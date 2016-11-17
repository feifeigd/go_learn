package main

import "fmt"
import "os"
import "encoding/gob"

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

func (p *Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, e := range p.Email {
		s += "\n" + e.Kind + ":" + e.Address
	}
	return s
}

func main() {
	var person Person
	loadGob("../SaveGob/person.gob", &person)
	fmt.Println("Person:", person.String())
}

func loadGob(file string, key interface{}) {
	in, err := os.Open(file)
	checkError(err)
	defer in.Close()
	decoder := gob.NewDecoder(in)
	err = decoder.Decode(key)
	checkError(err)
}

func checkError(err error) {
	if nil != err {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}
