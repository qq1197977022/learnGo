package demo10_UDP通信

import (
	"bytes"
	"fmt"
	"log"
	"net"
)

func udpServer() {
	defer wg.Done()

	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:7793")
	if err != nil {
		log.Fatalln(err)
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalln(err)
	}
	defer udpConn.Close()

	recBytes := make([]byte, 300)
	i, clientAddr, err := udpConn.ReadFromUDP(recBytes)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("server收到Client(%s)字节: %d\n", clientAddr, i)

	buffer := bytes.NewBuffer(recBytes[0:i])
	s := buffer.String()
	fmt.Printf("Client(%s)数据: %s\n", clientAddr, s)

	transpondUdpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	n, err := udpConn.WriteToUDP(recBytes[0:i], transpondUdpAddr) //服务端转发客户端数据至另一UDP地址
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Server转发字节: %d\n", n)

}
