package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	flag int32
	wg1  sync.WaitGroup
)

func main() {
	wg1.Add(2)

	go incCounter1("A")
	go incCounter1("B")

	time.Sleep(2 * time.Second)

	atomic.StoreInt32(&flag, 1) //写

	wg1.Wait()
}

func incCounter1(prefix string) {
	defer wg1.Done()
	for {
		fmt.Printf("--->%s\n", prefix)
		if atomic.LoadInt32(&flag) == 1 { //读
			break
		}
		time.Sleep(time.Second)
	}
}
