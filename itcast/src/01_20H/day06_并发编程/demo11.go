package main

import (
	"fmt"
	"time"
)

func main() {
	ch3 := make(chan int)
	select {
	case num := <-ch3:
		fmt.Printf("从ch接收 %d\n", num)
	case <-time.After(time.Second):
		fmt.Println("超时")
	}
}
