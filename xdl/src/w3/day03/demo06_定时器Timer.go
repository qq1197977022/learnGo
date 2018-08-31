package day03

import (
	"fmt"
	"time"
)

func Timer() {
	wg.Add(1)

	timer()

	wg.Wait()
}
func timer() { //创建定时器
	fmt.Println(time.Now())
	timerPtr1 := time.NewTimer(time.Second * 2) //方式一
	fmt.Println(<-timerPtr1.C)

	time.AfterFunc(time.Second, callBackFun) //方式二
}
func callBackFun() {
	defer wg.Done()

	fmt.Println("定时时间到...在当前goroutine调用回调函数")
}
