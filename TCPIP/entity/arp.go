package entity

type ArpFrame struct {
	HardwareType  []byte
	ProtocolType  []byte
	HardewareSize []byte
	ProtocolSize  []byte
	Opcode        []byte
	SenderMacAddr []byte
	SenderIpAddr  []byte
	TargetMacAddr []byte
	TargetIpAddr  []byte
}

// ARP Frame Payload (28 bytes)
// +---------------------+---------------------+---------------------+
// | Hardware Type (2B)  | Protocol Type (2B) | Hardware Size (1B)  |
// +---------------------+---------------------+---------------------+
// | Protocol Size (1B)  | Opcode (2B)        | Sender MAC (6B)     |
// +---------------------+---------------------+---------------------+
// | Sender MAC (cont.)  | Sender IP (4B)     | Target MAC (6B)     |
// +---------------------+---------------------+---------------------+
// | Target MAC (cont.)  | Target IP (4B)                              |
// +---------------------+---------------------+---------------------+

// ARP Request
// || ff:ff:ff:ff:ff:ff || 00:1a:2b:3c:4d:5e || 08:06 ||
// || 00:01 || 08:00 || 06 || 04 || 00:01 ||
// || 00:1a:2b:3c:4d:5e || 192.168.1.10 || 00:00:00:00:00:00 || 192.168.1.1 ||

// ARP Reply
// || 00:1a:2b:3c:4d:5e || 00:1b:4d:6f:7g:8h || 08:06 ||
// || 00:01 || 08:00 || 06 || 04 || 00:02 ||
// || 00:1b:4d:6f:7g:8h || 192.168.1.1 || 00:1a:2b:3c:4d:5e || 192.168.1.10 ||
