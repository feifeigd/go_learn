package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

// exe www.qq.com:80
func main() {
	if 2 != len(os.Args) {
		fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	//resule,err:=readFully(conn)
	result, err := readFully(conn)
	checkError(err)

	fmt.Println(string(result))
}

func checkError(err error) {
	if nil != err {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if nil != err {
			if io.EOF == err {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
