package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CD  = "CD"
	DIR = "DIR"
	PWD = "PWD"
)

func main() {
	listener, err := net.Listen("tcp", ":1202")
	checkError(err)
	for {
		conn, err := listener.Accept()
		if nil != err {
			continue
		}
		go handleClient(conn)
	}
}

func checkError(err error) {
	if nil != err {
		fmt.Println("Fatal error", err.Error())
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
		s := string(buf[0:n])
		fmt.Println("recv data:", s)
		// 解码
		if CD == s[0:2] {
			chdir(conn, s[3:])
		} else if DIR == s[0:3] {
			dirList(conn)
		} else if PWD == s[0:3] {
			pwd(conn)
		}
	}
}

func chdir(conn net.Conn, s string) {
	if nil == os.Chdir(s) {
		conn.Write([]byte("OK"))
		return
	}
	conn.Write([]byte("ERROR"))
}

func dirList(conn net.Conn) {
	defer conn.Write([]byte("\r\n"))
	dir, err := os.Open(".")
	if nil != err {
		return
	}
	names, err := dir.Readdirnames(-1)
	if nil != err {
		return
	}
	for _, nm := range names {
		conn.Write([]byte(nm + "\r\n"))
	}
}

func pwd(conn net.Conn) {
	s, err := os.Getwd()
	if nil != err {
		conn.Write([]byte(""))
		return
	}
	conn.Write([]byte(s))
}
