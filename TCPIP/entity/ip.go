package entity

type IPHeader struct {
	Version       []byte
	HeaderLength  []byte
	TypeOfService []byte
	PacketLength  []byte
	PacketID      []byte
	// FlagはFlags(1bit)とFragmentOffset(13bit)を合わせたもの
	Flag      []byte
	TTL       []byte
	Protocol  []byte
	Checksum  []byte
	SrcIpAddr []byte
	DstIpAddr []byte
}
