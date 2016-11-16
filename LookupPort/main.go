package main

import (
	. "fmt"
	"net"
	"os"
)

// exe tcp telnet
func main() {
	if 3 != len(os.Args) {
		Fprintf(os.Stderr, "Usage:%s network-type service\n", os.Args[0])
		os.Exit(1)
	}
	networkType, service := os.Args[1], os.Args[2]
	port, err := net.LookupPort(networkType, service)
	if nil != err {
		Println("Error:", err.Error())
		os.Exit(2)
	}
	Println("Service port", port)
}
