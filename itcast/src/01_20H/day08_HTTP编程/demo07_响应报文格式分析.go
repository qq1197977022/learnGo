package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"reflect"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	wg.Add(2)

	go server() //服务端

	time.Sleep(time.Second * 3) //客户端延迟启动
	go client()                 //客户端

	wg.Wait()
}
func client() {
	defer wg.Done()

	fmt.Println("客户端启动...")

	lTcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln(err)
	}
	rTcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalln(err)
	}
	tcpConn, err := net.DialTCP("tcp", lTcpAddr, rTcpAddr)
	defer tcpConn.Close()
	if err != nil {
		log.Fatalln(err)
	}

	reqHeader := "GET /go HTTP/1.1\r\n" +
		"Host: 127.0.0.1:8000\r\n" +
		"Connection: keep-alive\r\n" +
		"Upgrade-Insecure-Requests: 1\r\n" +
		"User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.62 Safari/537.36\r\n" +
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,* /*;q=0.8\r\n" +
		"Accept-Encoding: gzip, deflate, br\r\n" +
		"Accept-Language: en-US,en;q=0.9\r\n" +
		"\r\n"

	buf := bytes.NewBufferString(reqHeader)
	reqHeaderByts := buf.Bytes()
	tcpConn.Write(reqHeaderByts)

	recByts := make([]byte, 1024)
	n, err := tcpConn.Read(recByts)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s", recByts[:n])
}

func server() {
	defer wg.Done()

	fmt.Println("服务端启动...")
	http.HandleFunc("/go", handlerFun) //在DefaultServerMux中, 为当前模式注册处理器函数
	http.ListenAndServe("127.0.0.1:8000", nil)
	/*
		1.监听并通过调用handler来启动一个服务处理请求
		2.如果handler为nil, 使用DefaultServeMux
	*/
}
func handlerFun(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HelloWorld")

	fmt.Println("-------------------")
	fmt.Println(reflect.TypeOf(w))
	fmt.Println(reflect.TypeOf(w).Kind())
	fmt.Println("-------------------")
}

/*
一.响应报文
	响应行:	HTTP/1.1 200 OK
	响应头	Date: Mon, 04 Jun 2018 09:53:50 GMT
			Content-Length: 11
			Content-Type: text/plain; charset=utf-8
	头体空行
	响应体	HelloWorld

	-------------------------------------------------

	响应行	HTTP/1.1 404 Not Found
	响应头	Content-Type: text/plain; charset=utf-8
			X-Content-Type-Options: nosniff
			Date: Mon, 04 Jun 2018 09:56:53 GMT
			Content-Length: 19
	头体空行
	响应体	404 page not found
*/
