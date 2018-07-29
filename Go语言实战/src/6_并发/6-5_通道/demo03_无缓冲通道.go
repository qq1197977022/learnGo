package main

import (
	"fmt"
	"sync"
	"time"
)

var wg4 sync.WaitGroup //wait for a collection goroutine to finish

func main() {
	wg4.Add(1) //WaitGroup计数+1, main函数等待最后一位参赛者(goroutine)跑步结束

	ch := make(chan int) //创建整型无缓冲通道, 返回T而不是*T

	go run(ch) //创建goroutine

	ch <- 1 //往通道发送数据

	wg4.Wait() //阻塞, 直到WaitGroup计数=0, 即所有goroutine完成
}
func run(ch chan int) {
	var newRunner int

	runner := <-ch

	fmt.Printf("选手 %d 领跑...\n", runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("选手 %d 上场准备--->\n", newRunner)
		go run(ch) //创建goroutine, 但此时通道为空, 因而会在阻塞, 模拟选手上场准备, 等待接棒
	}

	time.Sleep(2 * time.Second) //跑步时间

	if runner == 4 {
		fmt.Printf("选手 %d 到达终点, 比赛结束", runner)
		wg4.Done()
	} else { //接力棒交接
		fmt.Printf("选手 %d --->交接接力棒---> %d\n", runner, newRunner)
		ch <- newRunner
	}
}
