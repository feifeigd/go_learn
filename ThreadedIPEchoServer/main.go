package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	service := ":1201"
	listener, err := net.Listen("tcp", service)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if nil != err {
			continue
		}
		go handleClient(conn) // run as a goroutine
	}
}

func checkError(err error) {
	if nil != err {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if nil != err {
			return
		}
		// 原样发回去
		fmt.Println(string(buf[0 : n-1]))
		_, err2 := conn.Write(buf[0:n])
		if nil != err2 {
			return
		}
	}
}
