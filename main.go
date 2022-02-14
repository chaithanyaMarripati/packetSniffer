package main

import (
	"fmt"
	"sync"

	"github.com/chaithanyaMarripati/packetSniffer/logger"
	"github.com/google/gopacket/pcap"
)

func main() {
	fmt.Println("Hello there, welcome to packet sniffer")
	interfaces, err := pcap.FindAllDevs()
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	for _, iFace := range interfaces {
		handle, err := pcap.OpenLive(iFace.Name, 65536, true, pcap.BlockForever)
		if err != nil {
			panic(err)
		}
		wg.Add(1)
		go logger.Log(handle, &wg)
	}
	wg.Wait()
}
