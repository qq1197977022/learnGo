package main

func main() {
	ch := make(chan string, 10)
	data := "abc"
	ch <- data
	data = <-ch
}

/*
1.保证共享资源的安全访问 and 消除竞争状态
	1.原子函数: 原子操作
	2.互斥锁: 临界区
	3.通道: 通过收发数据, 在goroutine间做同步
		1.在goroutine间架起管道, 并提供确保同步交换数据的机制
		2.声明通道时需要指定共享资源数据类型
			1.可通过通道共享的数据类型
				1.内置类型: 值/指针
				2.命名类型: 值/指针
				3.结构类型: 值/指针
				4.引用类型: 值/指针
	4.接收操作符 和 发送语句
		1.接收操作符(一元运算符): 从通道接收数据 	~ 	data := <-ch
		2.发送语句: 发送数据到通道   	~ 	ch <- data
*/
