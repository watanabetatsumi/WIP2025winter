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
		// 8bitのunsinged int（0～255）に変換
		i, _ := strconv.ParseUint(v, 10, 8)
		// 明示的に0,1で表現
		ipbyte = append(ipbyte, byte(i))
	}

	return ipbyte
}

// CPU標準のリトルエディアンから、ビックエディアン（ネットワークバイトオーダー）に変換
// 0x1234 -> 0x3412
func htons(i uint16) uint16 {
	// ２バイト長の数字の下位半分を左シフトして取り出す。-> 0xff00（マスク）を掛けて(AND演算)、低位バイトをゼロクリア
	// 上位半分を右シフトして取り出した後、OR演算
	return (i<<8)&0xff00 | i>>8
}
