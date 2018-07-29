package main

import (
	"fmt"
	"time"
)

func main() {
	go printer("ABCDEFGHIJKLMNOPQRSTUVWXYZ") //用户A
	go printer("abcdefghijklmnopqrstuvwxyz") //用户B

	time.Sleep(time.Second * 5)
}

func printer(s string) { //模拟打印机
	for i, v := range s {
		fmt.Printf("%d\t%c\n", i, v)
		time.Sleep(time.Second * 1) //打印机工作延时
	}
}
