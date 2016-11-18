package main

import (
	"fmt"
	"net/http"
	"os"
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
	fmt.Println(response.Status)
	for k, v := range response.Header {
		fmt.Println(k+":", v)
	}
}
