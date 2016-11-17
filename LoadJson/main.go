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
	var person Person
	loadJSON("../SaveJson/person.json", &person)
	fmt.Println("Person:", person.String())
}

func loadJSON(file string, key interface{}) {
	in, err := os.Open(file)
	checkError(err)
	defer in.Close()
	decoder := json.NewDecoder(in)
	err = decoder.Decode(key)
	checkError(err)
}

func checkError(err error) {
	if nil != err {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func (p *Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ":" + v.Address

	}
	return s
}
