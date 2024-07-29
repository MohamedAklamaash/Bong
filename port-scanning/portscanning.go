package main

import (
	"fmt"
	"net"
)

func PortScannerInRange() {
	for i := 1; i < 1024; i++ {
		addr := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf("Port Opened:%d", i)
	}
}
