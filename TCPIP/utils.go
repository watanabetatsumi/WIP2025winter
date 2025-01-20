package tcpip

import (
	"strconv"
	"strings"
)

type LocalIpMacAddr struct {
	LocalMacAddr []byte
	LocalIpAddr  []byte
	Index        int
}

// ##.##.##.##でipアドレスが入力されるのに対し、
// "."でパースして、10進数の数字を8ビット長のバイトに変換
func Ip2Byte(ip string) []byte {
	var ipbyte []byte

	for _, v := range strings.Split(ip, ".") {
		i, _ := strconv.ParseUint(v, 10, 8)
		ipbyte = append(ipbyte, byte(i))
	}

	return ipbyte
}
