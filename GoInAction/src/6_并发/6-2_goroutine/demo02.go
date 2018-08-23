package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg1 sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1) //设置GO运行时调度器可用逻辑CPU数: 把该逻辑CPU分配到某个物理CPU: 给每个物理CPU分配一个逻辑CPU runtime.GOMAXPROCS(runtime.NumCPU())

	wg1.Add(2) //设置WaitGroup计数=2

	go printPrime1("A")
	go printPrime1("B")

	wg1.Wait() //阻塞, 直到WaitGroup计数=0, 即所有goroutine完成, 避免goroutine在在有机会执行前main函数返回
}

func printPrime1(prefix string) {
	defer wg1.Done() //WaitGroup计数-1
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s---> %d\n", prefix, outer)
	}
	fmt.Println(prefix, "-------------完成-------------")
}
