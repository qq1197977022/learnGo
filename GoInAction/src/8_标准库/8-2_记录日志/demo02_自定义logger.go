package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger //记录所有日志
	Info    *log.Logger //重要信息
	Warning *log.Logger //警告
	Error   *log.Logger //错误
)

func init() {
	file, err := os.OpenFile("demo02_logger.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) //如果不存在则创建, 只写, 尾追加
	if err != nil {
		log.Fatalln("Failed to open error log tempFile: ", err)
	}

	Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile) //创建自定义Logger
	/*
		1.ioutil.Discard: io.Writer = devNull(0)
			1.所有Write调用都成功返回, 但不会做任何事
			2.即忽略所有写入的数据, 当某等级日志不重要时, 使用Discard变量可以禁用该等级日志
		2.Stdin, Stdout, Stderr均是已打开的文件, 分别指向标准输入/出/错误的文件描述符
			var (
				Stdin = NewFile(unitptr(syscall.Stdin), "dev/stdin")
				Stdout = NewFile(unitptr(syscall.Stdout), "dev/stdout")
				Stderr = NewFile(unitptr(syscall.Stderr), "dev/stderr")
			NewFile函数以文件描述符和名字作为参数返回新的*File
		)
	*/
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
func main() {
	Trace.Println("I have something stand to say")
	Info.Println("Special Infomation")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}
