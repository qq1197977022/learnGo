package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int32
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2) //设置WaitGroup计数=2

	go incCounter("A")
	go incCounter("B")

	wg.Wait() //阻塞, 直到WaitGroup计数=0, 即所有goroutine完成, 避免在goroutine执行前main函数就返回

	fmt.Println(counter)
}

func incCounter(prefix string) {
	defer wg.Done() //WaitGroup计数-1

	for i := 0; i < 2; i++ {
		//原子函数原子式操作counter, 并返回新值
		ret := atomic.AddInt32(&counter, 1) //原子函数AddInt32原子式操作对counter +1
		fmt.Printf("%s ---> couter = %d\n", prefix, ret)

		runtime.Gosched() //yield处理器, 让其他goroutine运行, 当前goroutine从当前线程退出, 退回到运行队列		~~~>	强制GO运行时调度器切换goroutine
	}
}

/*
1.原子函数会将同一时刻同时访问同一共享资源的goroutine同步, 即同一时刻只有一个goroutine访问共享变量
*/
