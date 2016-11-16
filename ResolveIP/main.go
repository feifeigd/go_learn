package main

import (
	. "fmt"
	"net"
	"os"
)

/// 查ip地址
func main() {
	if 2 != len(os.Args) {
		Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		Println("Usage:", os.Args[0], "hostname")
		os.Exit(1)
	}
	name := os.Args[1]
	addr, err := net.ResolveIPAddr("ip", name)
	if nil != err {
		Println("Resolution error", err.Error())
		os.Exit(1)
	}
	Println("Resolved address is ", addr.String())
}
