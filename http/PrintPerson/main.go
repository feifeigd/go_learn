package main

import (
	"html/template"
	"os"
)

type Job struct {
	Employer string
	Role     string
}

type Person struct {
	Name   string
	Age    int
	Emails []string
	Jobs   []*Job
}

const templ = `
The name is {{.Name}}.
The age is {{.Age}}.

{{range.Emails}}
	An email is {{.}}
{{end}}

{{with.Jobs}}
	{{range .}}
		An employer is {{.Employer}}
		and the role is {{.Role}}
	{{end}}
{{end}}
`

func main() {
	job1 := Job{"Monash", "Honorary"}
	job2 := Job{"Box Hill", "Head of HE"}
	person := Person{
		Name:   "jan",
		Age:    50,
		Emails: []string{"feifeigd@21cn.com", "505507456@qq.com"},
		Jobs:   []*Job{&job1, &job2},
	}
	//t := template.New("Person template")
	//t, err := t.Parse(templ)
	t, err := template.ParseFiles("Person.tpl")
	checkError(err)
	err = t.Execute(os.Stdout, person)
	checkError(err)
}

func checkError(err error) {

}
