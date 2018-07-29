package main

import (
	"fmt"
	"time"
)

func main() {
	//demand1()
	//demand2()
	demand3()
	fmt.Println("休眠结束")
}

//延时方式一: 休眠
func demand1() {
	time.Sleep(time.Second)
}

//延时方式二: Timer
func demand2() {
	timer := time.NewTimer(time.Second)
	<-timer.C

	timer.Stop()
}

//延时方式三: After ~ 和Timer等价
func demand3() {
	ch2 := time.After(time.Second)
	<-ch2
}
