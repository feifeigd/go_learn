package main

import (
	. "fmt"
	"net"
	"os"
)

func main() {
	if 2 != len(os.Args) {
		Fprintf(os.Stderr, "Usage: %s dotted-ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	dotAddr := os.Args[1]
	addr := net.ParseIP(dotAddr)

	if nil == addr {
		Println("Invalid address")
		os.Exit(1)
	}
	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	Println("The address is ", addr.String(), "\nDefault mask length is ", bits,
		"\nLeading ones count is ", ones,
		"\nMask is(hex) ", mask.String(),
		"\nNetwork is ", network.String())
}
