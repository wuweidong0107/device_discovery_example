package multicast

import (
	"log"
	"net"
)

const (
	maxDatagramSize = 8192
)

func Listen(address string, handler func(*net.UDPAddr, int, []byte)) {
	addr, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenMulticastUDP("udp4", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	conn.SetReadBuffer(maxDatagramSize)

	for {
		buffer := make([]byte, maxDatagramSize)
		numByte, src, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}

		handler(src, numByte, buffer)
	}
}
