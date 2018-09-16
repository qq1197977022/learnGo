package demo10_UDP通信

import (
	"sync"
	"time"
)

var wg sync.WaitGroup

func UDPConn() {
	wg.Add(2)

	go udpServer()
	time.Sleep(time.Second * 3)
	go udpClient()

	wg.Wait()
}
