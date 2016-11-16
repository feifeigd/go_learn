package main

import (
	. "fmt"
	"io/ioutil"
	"net"
	"os"
)

// exe www.qq.com:80
func main() {
	if 2 != len(os.Args) {
		Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	//resule,err:=readFully(conn)
	result, err := ioutil.ReadAll(conn)
	checkError(err)

	Println(string(result))
}

func checkError(err error) {
	if nil != err {
		Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}
