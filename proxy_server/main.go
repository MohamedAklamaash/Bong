package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// things that are done
	// listen for proxy request
	// handle that request (redirect the url somewhere else)
	listenaddr := ":1001"
	listener, err := net.Listen("tcp", listenaddr)
	if err != nil {
		fmt.Print(err)
		log.Fatal("Unable to bind to the port")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Unable to accept connection")
		}
		var targetAddr string = "http://www.youtube.com"
		go HandleConnection(conn, targetAddr)
	}
}

func HandleConnection(src net.Conn, targetAddr string) {
	// connect to the target server
	dst, err := net.Dial("tcp", targetAddr)
	if err != nil {
		log.Fatal("error in connecting to the target server")
	}
	defer dst.Close()
	// copy the response from the src obj to the destination
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatal(err)
		}
	}()
	// copy the response from target to src obj
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatal(err)
	}
}
