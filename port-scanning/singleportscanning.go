package main

import (
	"log"
	"net"
)

func PortScanning() {
	conn, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil { // err is nil means connection is established
		log.Printf("Port scanned: %s\n", conn.RemoteAddr().String())
		conn.Close()
	} else {
		log.Printf("Failed to scan port: %s\n", err.Error())
	}
}
