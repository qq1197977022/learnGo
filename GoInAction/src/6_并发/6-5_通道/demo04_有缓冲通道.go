package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//使用确定数量的goroutine来处理确定数量的工作
const (
	grNum    = 4  // goroutine数
	taskLoad = 10 // 工作量
)

var wg5 sync.WaitGroup //wait for a collection goroutine to finish

func init() {
	rand.Seed(time.Now().Unix()) //使用指定的种子值, 初始化默认资源到确定状态
}

func main() {
	tasks := make(chan string, taskLoad) //带缓冲字符串型通道

	wg5.Add(grNum) //设置WaitGroup计数=2

	for gr := 1; gr <= grNum; gr++ { //批量创建goroutine
		go worker(tasks, gr)
	}

	//发送数据到通道
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	close(tasks) //关闭通道, 不再发送数据到通道, 从通道接受完缓冲数据后, 若继续接收将返回通道类型的零值而不是阻塞

	wg5.Wait() //阻塞, 直到WaitGroup计数=0, 即所有goroutine完成
}

func worker(tasks chan string, worker int) {
	defer wg5.Done() //WaitGroup计数-1

	for {
		// 等待分配工作
		task, ok := <-tasks
		if !ok { //通道被关闭且其缓冲区数据被接收完
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			break
		}

		fmt.Printf("Worker: %d : Started %s\n", worker, task) //开工

		time.Sleep(time.Duration(rand.Int63n(100)) * time.Millisecond) //模拟工作时间

		fmt.Printf("Worker: %d : Completed %s\n", worker, task) //收工
	}
}

/*
1.带缓冲通道
	1.从通道接收数据前能存储数据
	2.不强制要求收发数据的goroutine必须同时准备好/完成
	3.只有通道内可用缓冲区 < 发送数据时, 发送goroutine才会阻塞 | 只有通道中没有数据时, 接收goroutine才会阻塞
*/
