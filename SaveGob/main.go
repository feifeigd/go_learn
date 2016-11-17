package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

type Person struct {
	Name  Name
	Email []Email
}

func main() {
	person := Person{
		Name{"Newmarch", "Jan"},
		[]Email{{"home", "jan@newmarch.name"}, {"work", "j.newmarch@boxhill.edu.au"}},
	}
	saveGob("person.gob", person)
}

func saveGob(file string, key interface{}) {
	out, err := os.Create(file)
	checkError(err)
	defer out.Close()
	encoder := gob.NewEncoder(out)
	err = encoder.Encode(key)
	checkError(err)
}

func checkError(err error) {
	if nil != err {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}
