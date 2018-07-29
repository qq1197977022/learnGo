package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	engine()
}

func engine() {
	tcpAd, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080") //解析TCP地址
	if err != nil {
		log.Fatalln(err)
	}
	tcpLisner, err := net.ListenTCP(tcpAd.Network(), tcpAd) //TCP监听
	if err != nil {
		log.Fatalln(err)
	}

	for {
		tcpCon, err := tcpLisner.AcceptTCP() //接受TCP请求
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%s连接成功\n", tcpCon.RemoteAddr().String())

		go handle(tcpCon)
	}
}

func handle(tcpCon *net.TCPConn) {
	senRecByts := make([]byte, 1024) //收发数据缓冲区
	go receive(tcpCon, senRecByts)
	go send(senRecByts, tcpCon)
}

func receive(tcpCon *net.TCPConn, senRecByts []byte) {
	for {
		length, err := tcpCon.Read(senRecByts) //接收
		if err != nil || err == io.EOF {
			log.Printf("%s\t客户端%s关闭连接", err, tcpCon.RemoteAddr().String())
			runtime.Goexit() //终止当前goroutine
		}
		fmt.Printf("接收: %s\n", senRecByts[:length])
	}
}

func send(senRecByts []byte, tcpCon *net.TCPConn) {
	for {
		fmt.Print("请输入: ")
		fmt.Scan(&senRecByts)
		length, err := tcpCon.Write(senRecByts) //发送
		if err != nil {
			log.Println(err)
			runtime.Goexit()
		}
		fmt.Printf("发送字节: %d\n", length)
	}
}
