/*MIT License

Copyright (c) 2022 betwowt

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"syscall"
	"time"
)

// 传入模版文件进行创建
// ex: md-template --templatePath=xxx --title "linux 应该如何学"

func main() {

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Println()
		fmt.Println("mdt --templatePath=xxx --title \"linux 应该如何学\" \n")
		flag.PrintDefaults()
	}

	var templatePath, title string
	flag.StringVar(&templatePath, "templatePath", "", "模版文件路径")
	flag.StringVar(&title, "title", "", "标题")
	flag.Parse()
	fmt.Println(templatePath, title)
	date := time.Now().Format("2006-01-02 15:04:05-0700")
	if len(templatePath) == 0 || len(title) == 0 {
		flag.Usage()
		syscall.Exit(1)
	}

	// 文件名是将title 转换成 英文+数字+-的格式
	filename := title + ".md"

	file, err := os.Open(templatePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	line := bufio.NewReader(file)

	wfile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wfile.Close()

	for {
		content, _, err := line.ReadLine()
		if err == io.EOF {
			break
		}

		fileContext := string(content)
		//fmt.Println(fileContext)
		fileContext = strings.ReplaceAll(fileContext, "$date", date)
		fileContext = strings.ReplaceAll(fileContext, "$title", title)

		fileContext = fileContext + "\n"
		wfile.Write([]byte(fileContext))
	}
}
