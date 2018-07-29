package main

import (
	"fmt"
	"log"
	"time"
)

var ch = make(chan bool)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	go personA()
	go personB()

	time.Sleep(time.Minute)
}

func personA() { //模拟用户A
	printer1("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	ch <- true
}

func personB() { //模拟用户B
	<-ch
	/*
		1.阻塞, 直到通道另一端准备好
		2.从通道接收并抛弃数据 ~ 这里仅仅是利用通道同步数据交互特性, 实现对打印机的同步访问
	*/
	printer1("abcdefghijklmnopqrstuvwxyz")
}

func printer1(s string) { //模拟打印机
	for i, v := range s {
		fmt.Printf("%d\t%c\n", i, v)
		time.Sleep(time.Second * 1) //打印机工作延时
	}
}
