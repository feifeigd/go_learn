package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
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

func main() {
	job1 := Job{"Monash", "Honorary"}
	job2 := Job{"Box Hill", "Head of HE"}
	person := Person{
		Name:   "jan",
		Age:    50,
		Emails: []string{"feifeigd@21cn.com", "505507456@qq.com"},
		Jobs:   []*Job{&job1, &job2},
	}
	tpl_name := "Person.tpl"
	tpls := template.New("")
	tpls.Funcs(template.FuncMap{"email_expand": EmailExpander})
	_, err := tpls.ParseGlob("templates/*.tpl")
	checkError(err)
	err = tpls.ExecuteTemplate(os.Stdout, tpl_name, person)
	checkError(err)
}

func checkError(err error) {
	if nil != err {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

// EmailExpander @转为 at
func EmailExpander(args ...interface{}) string {
	ok := false
	var s string
	if 1 == len(args) {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	// 找@
	substrs := strings.Split(s, "@")
	if 2 != len(substrs) {
		return s
	}
	return substrs[0] + " at " + substrs[1]
}
