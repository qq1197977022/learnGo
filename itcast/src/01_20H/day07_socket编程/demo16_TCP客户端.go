package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	engine1()
}

func engine1() {
	wg.Add(2)

	lTcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatalln(err)
	}
	rTcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln(err)
	}

	tcpCon, err := net.DialTCP(lTcpAddr.Network(), lTcpAddr, rTcpAddr) //TCP拨号连接
	defer tcpCon.Close()
	if err != nil {
		log.Fatalln(err)
	}

	senRecByts := make([]byte, 1024) //收发数据缓冲区
	go send1(tcpCon, senRecByts)
	go receive1(tcpCon, senRecByts)

	wg.Wait()
}

func receive1(tcpCon *net.TCPConn, senRecByts []byte) {
	for {
		length, err := tcpCon.Read(senRecByts) //接收
		if err != nil || err == io.EOF {
			log.Fatalf("%s\t服务端关闭连接", err)
		}
		fmt.Printf("接收: %s\n", senRecByts[:length])
	}
}

func send1(tcpCon *net.TCPConn, senRecByts []byte) {
	for {
		fmt.Print("请输入: ")
		fmt.Scan(&senRecByts)
		senRecByts = bytes.NewBuffer(senRecByts).Bytes()
		length, err := tcpCon.Write(senRecByts) //发送
		if err != nil {
			log.Println(err)
			runtime.Goexit()
		}
		fmt.Printf("发送字节: %d\n", length)
	}
}
