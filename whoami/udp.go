package main

import (
  "net"
  "log"
)

const MTU = 1500

func udpHandler(udpAddr *net.UDPAddr, msg []byte) ([]byte, error) {
	// echo
	return msg, nil
}

func udpServer(udpConn *net.UDPConn) {
	defer udpConn.Close()

	for {
		var buf = make([]byte, MTU)

		if read, udpAddr, err := udpConn.ReadFromUDP(buf); err != nil {
			log.Printf("ReadFromUDP: %v", err)
		} else if msg, err := udpHandler(udpAddr, buf[:read]); err != nil {
			log.Printf("UDP %v: %v", udpAddr, err)
		} else if _, err := udpConn.WriteTo(msg, udpAddr); err != nil {
			log.Printf("WriteToUDP %v: %v", udpAddr, err)
		} else {
			log.Printf("UDP %v", udpAddr)
		}
	}
}

func udpStart(listen string) {
  if udpAddr, err := net.ResolveUDPAddr("udp", listen); err != nil {
    log.Fatalf("ResolveUDPAddr %s: %v", listen, err)
  } else	if udpConn, err := net.ListenUDP("udp", udpAddr); err != nil {
    log.Fatalf("ListenUDP %v: %v", udpAddr, err)
  } else {
    go udpServer(udpConn)
  }
}
