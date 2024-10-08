package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var DevName = "wlp1s0"
var Found = false
var netInterfaces []string

func main() {
	devices, err := pcap.FindAllDevs() // finding network interfaces for packet capturing
	if err != nil {
		log.Fatal(err)
	}
	for _, device := range devices {
		if device.Name == DevName {
			Found = true
		}
		netInterfaces = append(netInterfaces, device.Name)
		fmt.Println(device)
	}
	if !Found {
		log.Fatal("desired device not found")
	}
	handle, err := pcap.OpenLive(DevName, 1600, false, pcap.BlockForever)
	if err != nil {
		log.Fatal("error in opening live")
	}
	defer handle.Close()
	// berkeley packet filter for filtering ftp traffic
	if err := handle.SetBPFFilter("tcp and port 443"); err != nil {
		log.Panicln(err)
	}
	source := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range source.Packets() {
		appLayer := packet.ApplicationLayer()
		if appLayer != nil {
			data := appLayer.Payload()

			if bytes.Contains(data, []byte("USER")) || bytes.Contains(data, []byte("PASS")) {
				fmt.Println(string(data))
			}
		}
		// fmt.Println(packet)
	}
}

func GetNetworkInterfaces() []string {
	return netInterfaces
}
