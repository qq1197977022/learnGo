package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	count int32
	wg    sync.WaitGroup //wait for a collection goroutine to finish
)

func main() {
	wg.Add(2) //设置WaitGroup计数=2

	go incCounter()
	go incCounter()

	wg.Wait() //阻塞, 等待WaitGroup计数=0, 即所有goroutine完成, 避免在goroutine有机会执行前main函数返回

	fmt.Println(count)
}

func incCounter() {
	defer wg.Done() //WaitGroup计数-1

	for i := 0; i < 2; i++ {
		value := count    //读
		runtime.Gosched() //yields处理器, 让其他goroutine运行, 当前goroutine从当前线程退出, 退回到运行队列		~~~>	强制GO运行时调度器切换goroutine
		value++
		count = value //写
	}
}

/*
1.竞争状态
	1.>=2个非同步goroutine同时访问(读写)同一共享资源
2.对同一共享资源的访问(读写)必须是原子化的: 同一时刻只能有一个goroutine访问(读写)
3.竞争状态检测
	1.步骤
		1.go build -race: 启用数据竞争检测
		2.运行 ---> 显示检测结果
4.消除竞争状态
	1.传统同步机制: 共享资源锁 ~ 顺序访问共享资源
	2.通信顺序进程(CSP)
*/
