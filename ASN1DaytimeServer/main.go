package main

import (
	"encoding/asn1"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":1200"
	listener, err := net.Listen("tcp", service)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if nil != err {
			continue
		}
		daytime := time.Now()
		mdata, _ := asn1.Marshal(daytime)
		conn.Write(mdata)
		conn.Close()
	}
}

func checkError(err error) {
	if nil != err {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}
