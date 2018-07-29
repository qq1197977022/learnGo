package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("当前时间: %v\n", time.Now())
	timer := time.NewTimer(time.Second * 2) //在指定时间后, 在其channel上发送当前时间
	fmt.Println(<-timer.C)
}
