package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	nBefore := runtime.GOMAXPROCS(1) //设置GO运行时调度器可的用最大逻辑CPU数: 把该逻辑CPU分配到某个物理CPU: 给每个物理CPU分配一个逻辑CPU runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(nBefore)             //之前逻辑CPU数
	fmt.Println(runtime.NumCPU())    //获取当前线程可用逻辑CPU数

	var wg sync.WaitGroup //wait for a collection goroutine to finish

	wg.Add(2) //WaitGrou计数+2

	fmt.Println("开始groutine")

	go func() { //匿名函数
		defer wg.Done() //WaitGroup计数-1

		//显示小写字母表3次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println() //换行
		}
	}()

	//显示大写字母表3次
	go func() {
		defer wg.Done() //WaitGroup计数-1

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println() //换行
		}
	}()

	wg.Wait() //阻塞, 直到WaitGrou计数=0, 即所有goroutine结束, 避免main函数在goroutine有机会执行前返回

	fmt.Println("程序结束!")
}

/*
1.goroutine创建并管理其寿命
2.基于调度器的内部算法, 正在运行的goroutine结束前, 可以被停止(放回运行队列)并重新调度, 以避免某goroutine长时间占用逻辑CPU
*/
