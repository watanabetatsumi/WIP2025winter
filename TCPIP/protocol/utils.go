package protocol

import (
	"fmt"
	"net"
	"reflect"
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

// 指定したインターフェースのMACアドレスとIPアドレスを返す
func getLocalIpAddr(ifname string) (localif LocalIpMacAddr, err error) {
	nif, err := net.InterfaceByName(ifname)
	if err != nil {
		return localif, err
	}

	localif.LocalMacAddr = nif.HardwareAddr
	localif.Index = nif.Index

	// ローカルのインターフェースのIPアドレスをすべて返す
	addrs, err := nif.Addrs()
	if err != nil {
		return localif, err
	}
	for _, addr := range addrs {
		// 無名関数で、型アサーションを行う -> ok(=1)なら、if文内の処理を行う
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			// To4()はIPアドレスを4byte長の0,1で表す関数
			if ipnet.IP.To4() != nil {
				localif.LocalIpAddr = ipnet.IP.To4()
				break
			}
		}
	}

	return localif, nil
}

// goのinterface{}型はオールマイティを指す->逆にValueOf関数を使って具体的な方を特定しなきゃいけない
func toByteArr(value interface{}) []byte {
	rv := reflect.ValueOf(value)
	// 構造体かどうかチェック
	if rv.Kind() != reflect.Struct {
		panic("toByteArr: value is not a struct")
	}

	var arr []byte

	// 構造体のfieldをひとつづつ取り出して、
	for i := 0; i < rv.NumField(); i++ {
		// 一旦interface{}型に変換して、[]byte型に変換
		byteO := rv.Field(i).Interface().([]byte)
		arr = append(arr, byteO...)
	}

	return arr
}

func printByteArr(arr []byte) string {
	var str string

	// %xは値を16進数の文字で表示する。
	for _, v := range arr {
		str += fmt.Sprintf("%x", v)
	}

	return str
}
