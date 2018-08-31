package day02

import (
	"fmt"
	"time"
)

func Goroutine1() {
	var counter int

	for {
		fmt.Println("------------main goroutine终止, 但main函数没有返回------------")
		time.Sleep(time.Second)

		counter++
		if counter == 5 {
			break
		}
	}
}
