package main

import (
	"log"
	"os"
)

func main() {
	err := os.RemoveAll("D:\\新建文本文档.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

/*
一.文件类型
	1.设备文件
		1.屏幕: 标准输出
		2.键盘: 标准输入
	2.磁盘文件
		1.文本文件
		2.二进制文件
二.为什么需要文件
	1.持久化存储数据
*/

//该Read了...................
