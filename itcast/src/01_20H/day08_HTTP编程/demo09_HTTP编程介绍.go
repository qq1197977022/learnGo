package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	wg1.Add(1)

	go server1()

	time.Sleep(time.Second * 3) //客户端延时启动
	go client1()

	wg1.Wait()
}

func server1() {
	defer wg1.Done()

	http.HandleFunc("/", handler1)
	http.ListenAndServe("127.0.0.1:8000", nil)
}
func handler1(w http.ResponseWriter, r *http.Request) {
	senByts := bytes.NewBufferString("HelloWorld").Bytes()
	w.Write(senByts)
}

func client1() {
	defer wg1.Done()

	res, err := http.Get("http://127.0.0.1:8000")
	if err != nil {
		log.Fatalln(err)
	}
	recByts := make([]byte, 16)
	n, err := res.Body.Read(recByts)
	defer res.Body.Close()
	if err != nil && err != io.EOF { //读取len(recByts)字节, 服务端发送数据太少 <len(recByts) 导致EOF
		log.Fatalln(err)
	}
	fmt.Printf("读取字节: %d\n", n)
	fmt.Printf("%s\n", recByts[:n])
}

//Go标准库http包提供了HTTP客户端和服务端实现
