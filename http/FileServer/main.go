package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fileServer := http.FileServer(http.Dir("./html"))
	http.Handle("/", fileServer)
	http.HandleFunc("/cgi-bin/printenv", printEnv)
	err := http.ListenAndServe(":8000", nil)
	checkError(err)
}

func checkError(err error) {
	if nil != err {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(4)
	}
}

func printEnv(w http.ResponseWriter, req *http.Request) {
	env := os.Environ()
	w.Write([]byte("<h1>Environment</h1>\n<pre>"))
	for _, v := range env {
		w.Write([]byte(v + "\n"))
	}
	w.Write([]byte("</pre>"))
}
