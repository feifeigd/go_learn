package main

import (
	"encoding/json"
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
		Name:  Name{"Newmarch", "Jan"},
		Email: []Email{{"home", "jan@newmarch.name"}, {"work", "j.newmarch@boxhill.edu.au"}},
	}
	saveJSON("person.json", person)
}

func saveJSON(file string, key interface{}) {
	out, err := os.Create(file)
	checkError(err)
	defer out.Close()
	encoder := json.NewEncoder(out)
	err = encoder.Encode(key)
	checkError(err)
}

func checkError(err error) {
	if nil != err {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
