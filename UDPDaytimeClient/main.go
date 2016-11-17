package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if 2 != len(os.Args) {
		fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	_, err = conn.Write([]byte("anything"))
	checkError(err)

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)
	fmt.Println(string(buf[0:n]))
}

func checkError(err error) {
	if nil != err {
		fmt.Fprint(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
