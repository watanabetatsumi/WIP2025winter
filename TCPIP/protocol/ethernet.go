package protocol

import "github.com/WIP2025winter/TCPIP/entity"

func NewEthernet(dstMacAddr []byte, srcMacAddr []byte, ethType string) entity.EthernetFrame {
	frame := entity.EthernetFrame{
		DstMACAddr: dstMacAddr,
		SrcMACAddr: srcMacAddr,
	}

	// イーサーネットフレームにはバイト列としてtypeを指定する。
	switch ethType {
	case "IPv4":
		frame.Type = entity.IPv4
	case "ARP":
		frame.Type = entity.ARP
	}

	return frame
}
