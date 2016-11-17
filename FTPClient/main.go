package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
)

// 用户输入命令
const (
	uiDir  = "dir"
	uiCd   = "cd"
	uiPwd  = "pwd"
	uiQuit = "quit"
)

// 发给服务器的指令
const (
	DIR = "DIR"
	CD  = "CD"
	PWD = "PWD"
)

func main() {
	if 2 != len(os.Args) {
		fmt.Println("Usage:", os.Args[0], " host")
		os.Exit(1)
	}
	host := os.Args[1]
	conn, err := net.Dial("tcp", host+":1202")
	checkError(err)
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if nil != err {
			break
		}
		line = strings.TrimRight(line, "\t\r\n") // 截去右边空白
		// command arg
		strs := strings.SplitN(line, " ", 2)
		fmt.Println("input:", strs)
		switch strs[0] {
		case uiDir:
			dirRequest(conn)
		case uiCd:
			if 2 != len(strs) {
				fmt.Println("cd <dir>")
				continue
			}
			fmt.Println("CD \"", strs[1], "\"")
			cdRequest(conn, strs[1])
		case uiPwd:
			pwdRequest(conn)
		case uiQuit:
			conn.Close()
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
		}
	}
}

func checkError(err error) {
	if nil != err {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}

func dirRequest(conn net.Conn) {
	conn.Write([]byte(DIR + " "))
	var buf [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, _ := conn.Read(buf[0:])
		result.Write(buf[0:n])
		length := result.Len()
		contents := result.Bytes()
		if "\r\n\r\n" == string(contents[length-4:]) {
			fmt.Println(string(contents[0 : length-4]))
			return
		}
	}
}

func cdRequest(conn net.Conn, dir string) {
	conn.Write([]byte(CD + " " + dir))
	var response [512]byte
	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	if "OK" != s {
		fmt.Println("Failed to change dir")
	}
}

func pwdRequest(conn net.Conn) {
	conn.Write([]byte(PWD))
	var response [512]byte
	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	fmt.Println("Current dir\"" + s + "\"")
}
