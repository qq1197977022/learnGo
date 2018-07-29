package main

import (
	"fmt"
	"sync"
	"time"
)

var ch2 = make(chan int, 3)
var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go productor()
	go consumer()

	wg.Wait()
}

func productor() {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		ch2 <- i
		fmt.Printf("生产%d\n", i)
	}
	close(ch2) //不再发送数据到ch
}

func consumer() {
	defer wg.Done()

	time.Sleep(time.Second * 5)
	for v := range ch2 {
		fmt.Printf("消费 %d\t\n", v)
	}
}
