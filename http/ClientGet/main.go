package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if 2 != len(os.Args) {
		fmt.Println("Usage:", os.Args[0], "http://host:port/page")
		os.Exit(1)
	}
	url, err := url.Parse(os.Args[1])
	checkError(err)
	client := &http.Client{}
	request, err := http.NewRequest("GET", url.String(), nil)
	checkError(err)
	// only accept UTF-8
	request.Header.Add("Accept-Charset", "UTF-8;q=1,ISO-8859-1;q=0")

	response, err := client.Do(request)
	if nil != err {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	if "200 OK" != response.Status {
		fmt.Println(response.Status)
		os.Exit(2)
	}
	chSet := getCharset(response)
	fmt.Printf("got charset %s\n", chSet)
	if "UTF-8" != chSet {
		fmt.Println("Cannot handle ", chSet)
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

func checkError(err error) {
	if nil != err {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}

func getCharset(response *http.Response) string {
	contentType := response.Header.Get("Content-Type")
	if "" == contentType {
		return "UTF-8" // 可能是UTF-8
	}
	idx := strings.Index(contentType, "charset:")
	if -1 == idx {
		return "UTF-8"
	}
	return strings.Trim(contentType[idx:], " ")
}
