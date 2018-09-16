package main

import (
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	//day01.DataFetchParse()

	/*
		go day02.Goroutine1()
		runtime.Goexit()

		1.main函数所在goroutine终止, 但没有返回, 所以其他goroutine正常执行
		2.若此时其他goroutine均终止而不是返回, 则程序崩溃
		3.即使其他goroutine正常返回程序依然崩溃*/

	//day03.CloseChan()
	//day03.Timer()
	//day03.Ticker()
	//demo10_UDP通信.UDPConn()
}
