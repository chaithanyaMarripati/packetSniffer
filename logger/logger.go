package logger

import (
	"fmt"
	"sync"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

//Logger will take pcap handle and read any data the goes through it
//will log the ip address for the decoded packet data
func Log(handle *pcap.Handle, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		data, _, err := handle.ReadPacketData()
		if err != nil {
			fmt.Println("couldn't read packet data")
		}
		packet := gopacket.NewPacket(data, layers.LayerTypeEthernet, gopacket.NoCopy)
		networkLayer := packet.NetworkLayer()
		if networkLayer == nil {
			continue
		}
		networkFlow := networkLayer.NetworkFlow()
		fmt.Println("packet moving from ", networkFlow.Src().String(), " to ", networkFlow.Dst().String())
	}
}
