package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 100000; i++ {
		}
		fmt.Println("hello")
	}()

	for i := 0; i < 3; i++ {
		fmt.Println("world")
		runtime.Gosched()
	}
}
