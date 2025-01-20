package entity

var (
	// 2byte(0x11=255)
	// 0x08が上位ビット。0x00が下位ビット。
	IPv4 = []byte{0x08, 0x00}
	ARP  = []byte{0x08, 0x06}
)

type EthernetFrame struct {
	// byteは1byteの型
	DstMACAddr []byte
	SrcMACAddr []byte
	Type       []byte
}
