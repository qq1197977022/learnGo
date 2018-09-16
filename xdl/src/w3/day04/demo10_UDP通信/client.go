package demo10_UDP通信

import (
	"bytes"
	"fmt"
	"log"
	"net"
)

func udpClient() {
	defer wg.Done()

	rUdpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:7793")
	if err != nil {
		log.Fatalln(err)
	}

	udpConn, err := net.DialUDP("udp", nil, rUdpAddr)
	if err != nil {
		log.Fatalln(err)
	}
	defer udpConn.Close()

	buffer := bytes.NewBufferString("aaaBBB")
	sendBytes := buffer.Bytes()
	n, err := udpConn.Write(sendBytes)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Client发送字节: %d\n", n)
}
