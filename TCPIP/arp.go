package tcpip

import "github.com/WIP2025winter/TCPIP/entity"

func NewArpRequest(local LocalIpMacAddr, targetIP string) entity.ArpFrame {
	arp := entity.ArpFrame{
		// イーサーネットの場合、0x0001(2バイト長)で固定。
		HardwareType: []byte{0x00, 0x01},
		// IPv4の場合。0x0800で固定。
		ProtocolType: []byte{0x08, 0x00},
		// MACアドレスのサイズ（バイト）。0x06(6バイト)
		HardewareSize: []byte{0x06},
		// IPアドレスのサイズ（バイト）。0x04(4バイト)
		ProtocolSize: []byte{0x04},
		// ARPリクエスト -> １(0x0001)
		Opcode: []byte{0x00, 0x01},
		// 送信元MACアドレス
		SenderMacAddr: local.LocalMacAddr,
		// 送信元IPアドレス
		SenderIpAddr: local.LocalIpAddr,
		// ターゲットMACアドレス -> ARP Requestなのでブロードキャスト=0
		TargetMacAddr: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		// ターゲットIPアドレス
		TargetIpAddr: Ip2Byte(targetIP),
	}

	return arp
}

func (*entity.ArpFrame) Send(dstMacAddr int, packet []byte) entity.ArpFrame {

}
