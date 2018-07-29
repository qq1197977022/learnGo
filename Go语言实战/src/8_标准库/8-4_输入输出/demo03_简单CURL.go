package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.baidu.com") //get请求
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create("demo03_curl.txt") //创建文件
	if err != nil {
		log.Fatalln(err)
	}

	writers := io.MultiWriter(os.Stdout, file) //创建一个writer, 重复写入到所有提供的writer
	io.Copy(writers, resp.Body)

	if err := resp.Body.Close(); err != nil { //关闭ReaderCloser
		log.Fatalln(err)
	}

	if err := file.Close(); err != nil { //关闭文件, 释放资源
		log.Println(err)
	}
}

/*
1.Linux和MacOS中有名为curl的命令行工具, 可以对指定URL发起HTTP请求, 并保存返回内容
*/
