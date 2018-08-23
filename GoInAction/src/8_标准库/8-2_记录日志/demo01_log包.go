package main

import (
	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	log.Println("message") //输出标准日志记录器

	log.Fatalln("fatal message") //log.Println() && os.Exit(1)

	log.Panicln("panic message") //log.Println() && log.panic()
}

/*
1.关键字iota在常量声明区中有特殊作用
	1.让编译器为每个常量赋值相同表达式, 直到声明结束 或 遇到一个新的赋值语句
	2.初始值=0, iota的值在每次处理为常量后, 都会自增1
	eg.
		const (
			Ldate = 1 << iota 	// 1 << 0 = 000000001 = 1
			Ltime 				// 1 << 1 = 000000010 = 2
			Lmicroseconds  		// 1 << 2 = 000000100 = 4
			Llongfile 			// 1 << 3 = 000001000 = 8
			Lshortfile 			// 1 << 4 = 000010000 = 16
			...
			)
2.log包实现了日志记录器多goroutine安全
	1.多goroutine同时调用同一个logger的函数而不会冲突
	2.预定义logger和自定义logger均满足该性质
*/
