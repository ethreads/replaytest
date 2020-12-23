package capture

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"log"
	"os"
	"time"
)

func capture()  {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(devices)

	var (
		device string = "en0"
		snapshotLen int32  = 1024
		timeout      time.Duration = 10 * time.Second
	)
	handle, err := pcap.OpenLive(device, snapshotLen, false, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	f, err := os.Create(fmt.Sprintf("%s.pcap", device))
	if err != nil {
		log.Fatal(err)
	}
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(uint32(snapshotLen), layers.LinkTypeEthernet)
	defer f.Close()
	for packet := range packetSource.Packets() {
		decodePacketInfo(packet)
		w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
	}
}

func decodePacketInfo(packet gopacket.Packet)  {
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
	}
}