package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i < 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			addr := fmt.Sprintf("localhost:%d", j)
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				// Port is closed or unreachable; nothing to log here
				return
			}
			defer conn.Close()
			fmt.Printf("Port Opened: %d\n", j)
		}(i)
	}
	wg.Wait()
}
