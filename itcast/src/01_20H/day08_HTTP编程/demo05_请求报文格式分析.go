package main

import (
	"fmt"
	"log"
	"net"
)

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalln(err)
	}
	tcpListener, err := net.ListenTCP(tcpAddr.Network(), tcpAddr)
	if err != nil {
		log.Fatalln(err)
	}
	tcpConn, err := tcpListener.AcceptTCP()
	if err != nil {
		log.Fatalln(err)
	}
	senRecByts := make([]byte, 1024)
	n, err := tcpConn.Read(senRecByts)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("接收字节数: %d\n", n)
	fmt.Printf("%s", senRecByts[:n])
}

/*
一.GET请求
	请求行		GET /?user=AAAAAAAAAAA&password=BBBBBBBBBBBBBBB HTTP/1.1
	请求头		Host: 127.0.0.1:8000
				Connection: keep-alive
				Upgrade-Insecure-Requests: 1
				User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.62 Safari/537.36
				Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,* /*;q=0.8
				Accept-Encoding: gzip, deflate, br
				Accept-Language: en-US,en;q=0.9
	头体空行
	请求体
二.Post请求
	请求行		POST / HTTP/1.1
	请求头		Host: 127.0.0.1:8000
				Connection: keep-alive
				Content-Length: 45
				Cache-Control: max-age=0
				Upgrade-Insecure-Requests: 1
				Origin: null
				Content-Type: application/x-www-form-urlencoded
				User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.62 Safari/537.36
				Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,* /*;q=0.8
				Accept-Encoding: gzip, deflate, br
				Accept-Language: en-US,en;q=0.9
	头体空行
	请求体		user=AAAAAAAAAAA&password=BBBBBBBBBBBBBBBBBBB
三.总结
	1.GET请求数据在请求行中, 请求体为空
	2.POST请求数据数据在请求体中
*/
