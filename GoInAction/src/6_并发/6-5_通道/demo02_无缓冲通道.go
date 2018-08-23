package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup //wait for a collection of goroutine to finish

func init() {
	rand.Seed(time.Now().UnixNano()) //使用指定种子值, 初始化默认资源到确定状态
}
func main() {
	ch := make(chan int) //创建int类型无缓冲通道

	wg.Add(2) //WaitGroup计数+2

	go player("A", ch)
	go player("B", ch)

	ch <- 1 //发球 ~~~> 相当于裁判把球抛给选手?

	wg.Wait() //阻塞, 直到WaitGroup计数=0, 即所有goroutine结束
}

//模拟选手
func player(name string, ch chan int) {
	defer wg.Done() //计数-1

	for {
		// 等待球被击打过来
		data, ok := <-ch

		/*1.yield一个额外的无类型布尔值, 检测通信是否成功
		1.true: 成功接收发送到通道的数据
		2.false: 因通道关闭或空而生成0值*/

		if !ok { // 如果通道被关闭或空, 我们就赢了
			fmt.Printf("恭喜 %s 获胜!\n", name)
			fmt.Printf("\n因通道关闭/空生而成0值---------------->%d", data)
			break
		}

		// 用特定随机数模拟丢球

		if rNum := rand.Intn(100); rNum%13 == 0 {
			// 丢球关闭通道, 表示我们输了
			close(ch)
			fmt.Printf("\n随机数rnum = %d\n\n", rNum)
			fmt.Printf("糟糕 %s 丢球!\n", name)
			break
		}

		// 显示击球次数, 并+1
		fmt.Printf("%s 第 %d 次击球\n", name, data)
		data++

		ch <- data //将球打向对手
	}
}

/*
1.无缓冲通道
	1.从通道接收数据前不能保存数据
	2.工作前提
		1.数据收发goroutine必须同时准备好, 才能完成数据收发操作
			1.任一方没有准备好, 会导致先准备好的goroutine阻塞等待(在通道中被锁住, 直到数据收发完成才被释放)
			2.数据收发过程两个goroutine都会在通道中被锁住, 直到数据收发完成
		2.这种通道收发数据的交互行为本身就是同步的 --- 任意一个操作都无法离开另一个操作独立存在
	3.数据收发过程
		1.负责数据收发的goroutine准备
		2.先准备好的goroutine在通道中被锁住, 同时准备好则两个goroutine都在通道中同时被锁住
		3.数据收发
		4.数据收发完成, 两个goroutine同时被释放
*/
