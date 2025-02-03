package tcpip

import (
	"log"
	"syscall"

	"github.com/WIP2025winter/TCPIP/entity"
)

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
	// syscall.SockaddrLinkLayerという形（RAWソケット向け）を介して、カーネル空間のSocketディスクリプタに
	// 書き込む。
	addr := syscall.SockaddrLinklayer{
		Protocol: syscall.ETH_P_ALL,    // すべてのプロトコル
		Ifindex:  dstMacAddr,           // どのネットワークインターフェースを使うか指定。（引数のマックアドレス）
		Hatype:   syscall.ARPHRD_ETHER, // イーサーネットフレーム
	}

	// "ソケット"を作成。->ソケットディスクリプタ（インターフェース）をオープン
	// ソケットディスクリプタをカーネル空間から持ち上げて、通信を開始（読み込み/書き込み）。
	// 通信の際のオプションは
	// syscall.AF_PACKET: 通信のレベルを指定。
	// syscall.SOCKRAW: デバイスドライバから受け取る。
	// syscall.プロトコルを指定。->すべて
	//  傍受したいのは、イーサーネットフレームのうちARP部分なのか、IPパケット部分なのか、ペイロード部分なのか等指定。
	sendfd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(htons(syscall.ETH_P_ALL)))
	if err != nil {
		log.Fatal("create sendfd err : %v\n", err)
	}
	defer syscall.Close(sendfd)

	// sendfdにpacketの内容を書き込んで、addr宛に送信(0はフラグ)
	// 書き込み専用ソケット
	err = syscall.Sendto(sendfd, packet, 0, &addr)
	if err != nil {
		log.Fatalf("Send to err : %v\n", err)
	}

	// レスポンスをListenするために待ち続ける -> forループ
	for {
		// ARPパケット自体は28バイト
		// イーサーネットとARPは同じL2レイヤ->気にする必要はない。
		recvBuf := make([]byte, 80)
		// sendfdディスクリプタ（インターフェース）を介して、レスポンスを読み取り、recvBufに書き込む
		_, _, err := syscall.Recvfrom(sendfd, recvBuf, 0)
		if err != nil {
			log.Fatalf("read err : %v\n", err)
		}

		if recvBuf[12] == 0x08 && recvBuf[13] == 0x06 {
			//  ARPのopcodeがReply(0x0002)かチェック
			if recvBuf[20] == 0x00 && recvBuf[21] == 0x02 {
				// 14バイト目以降のARPパケットのペイロードを読み込む。
				return parseArpPacket(recvBuf[14:])
			}
		}
	}
}

// byteオーダーを構造体にマッピングする関数
func parseArpPacket(packet []byte) entity.ArpFrame {
	return entity.ArpFrame{
		HardwareType:  []byte{packet[0], packet[1]},
		ProtocolType:  []byte{packet[2], packet[3]},
		HardewareSize: []byte{packet[4]},
		ProtocolSize:  []byte{packet[5]},
		Opcode:        []byte{packet[6]},
		SenderMacAddr: packet[8:14],
		SenderIpAddr:  packet[14:18],
		TargetMacAddr: packet[18:24],
		TargetIpAddr:  packet[24:28],
	}
}
