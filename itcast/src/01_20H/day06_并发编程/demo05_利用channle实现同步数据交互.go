package main

import "fmt"

var ch1 = make(chan string)

func main() {
	defer fmt.Println("main协程结束")

	go func() {
		for i, v := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
			fmt.Printf("%d\t%c", i, v)
		}
		ch1 <- "\n子协程结束"
	}()

	str := <-ch1
	fmt.Println(str)
}
