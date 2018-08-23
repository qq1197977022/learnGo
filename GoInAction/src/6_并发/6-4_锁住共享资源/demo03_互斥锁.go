package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	counter2 int32
	wg2      sync.WaitGroup
	mutex    sync.Mutex
)

func main() {
	wg2.Add(2)

	go incCounter2("A")
	go incCounter2("B")

	wg2.Wait()

	fmt.Println(counter2)
}

func incCounter2(prefix string) {
	defer wg2.Done()

	for i := 0; i < 2; i++ {
		mutex.Lock()
		{
			if prefix == "A" {
				time.Sleep(2 * time.Second)
			}
			value := counter2
			runtime.Gosched()
			/*
				1.yield处理器, 让其他goroutine运行, 当前goroutine从当前线程退出, 退回到运行队列		~~~>	强制GO运行时调度器切换goroutine
				2.但因为Mutex, 其他goroutine被锁定, GO运行时调度器会再次调度该goroutine运行
			*/
			value++
			counter2 = value

			fmt.Printf("%s--->%d\n", prefix, counter2)
		}
		mutex.Unlock()
	}
}

/*
1.互斥锁: 创建临界区, 同一时间只能有一个goroutine进入临界区(临界区同一时间只能有一个goroutine)
2.Mutex不和特定的goroutine, 任意goroutine锁定的Mutex, 允许任意goroutine解锁
*/
