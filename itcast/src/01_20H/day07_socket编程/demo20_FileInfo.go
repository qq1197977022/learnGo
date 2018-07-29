package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	stat()
	fmt.Println("=======================")
	lstat()
}

func stat() {
	//fInfo, err := os.Stat("D:\\workspace\\备忘.txt")	//原文件
	fInfo, err := os.Stat("C:\\Users\\kc\\Desktop\\备忘.lnk") //快捷方式
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Name: %s\n", fInfo.Name())
	fmt.Printf("Size: %d\n", fInfo.Size())
	fmt.Printf("Mode: %s\n", fInfo.Mode())
	fmt.Printf("ModTime: %s\n", fInfo.ModTime())
	fmt.Printf("IsDir: %t\n", fInfo.IsDir())
	fmt.Printf("Sys: %v\n", fInfo.Sys())
}

func lstat() {
	//fInfo, err := os.Lstat("D:\\workspace\\备忘.txt")	//原文件
	fInfo, err := os.Lstat("C:\\Users\\kc\\Desktop\\备忘.lnk") //快捷方式
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Name: %s\n", fInfo.Name())
	fmt.Printf("Size: %d\n", fInfo.Size())
	fmt.Printf("Mode: %s\n", fInfo.Mode())
	fmt.Printf("ModTime: %s\n", fInfo.ModTime())
	fmt.Printf("IsDir: %t\n", fInfo.IsDir())
	fmt.Printf("Sys: %v\n", fInfo.Sys())
}
