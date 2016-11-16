package main

import (
	. "fmt"
	"net"
	"os"
)

func main() {
	if 2 != len(os.Args) {
		Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addrs, err := net.LookupHost(name)
	if nil != err {
		Println("Error: ", err.Error())
		os.Exit(2)
	}
	for _, s := range addrs {
		Println(s)
	}
}
