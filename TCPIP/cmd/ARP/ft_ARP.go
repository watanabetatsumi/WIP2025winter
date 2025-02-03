package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/WIP2025winter/TCPIP/protocol"
)

func main() {

	var ipAddr string

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <ip address>")

		return
	}

	for i, v := range os.Args {
		if i == 1 {
			ipAddr = v
		}
	}

	ipaddr := strings.Split(ipAddr, ".")
	if len(ipaddr) != 4 {
		fmt.Println("Invalid IP address")

		return
	}

	protocol.ARP(ipAddr)
}
