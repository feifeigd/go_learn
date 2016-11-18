package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	if 2 != len(os.Args) {
		fmt.Println("Usage:", os.Args[0], "host:port")
		os.Exit(1)
	}
	url := os.Args[1]
	response, err := http.Head(url)
	if nil != err {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	if "200 OK" != response.Status {
		fmt.Println(response.Status)
		os.Exit(2)
	}
	b, _ := httputil.DumpResponse(response, false)
	fmt.Println(string(b))

	contentType := response.Header["Content-Type"]
	if !acceptableCharset(contentType) {
		fmt.Println("Cannot handle", contentType)
		os.Exit(4)
	}
	var buf [512]byte
	reader := response.Body
	for {
		n, err := reader.Read(buf[0:])
		if nil != err {
			os.Exit(0)
		}
		fmt.Print(string(buf[0:n]))
	}
}

func acceptableCharset(contentTypes []string) bool {
	// 仅支持UTF-8
	for _, cType := range contentTypes {
		if -1 != strings.Index(cType, "UTF-8") {
			return true
		}
	}
	return false
}
