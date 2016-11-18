package main

import (
	"errors"
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
	/*job1 := Job{"Monash", "Honorary"}
	job2 := Job{"Box Hill", "Head of HE"}
	person := Person{
		Name:   "jan",
		Age:    50,
		Emails: []string{"feifeigd@21cn.com", "505507456@qq.com"},
		Jobs:   []*Job{&job1, &job2},
	}*/
	tpl_name := "Person.tpl"
	tpls := template.New("")
	tpls.Funcs(template.FuncMap{"email_expand": EmailExpander, "sequence": sequenceFunc, "cycle": cycleFunc})
	_, err := tpls.ParseGlob("templates/*.tpl")
	checkError(err)
	//err = tpls.ExecuteTemplate(os.Stdout, tpl_name, person)
	err = tpls.ExecuteTemplate(os.Stdout, tpl_name, []string{"a", "b", "c", "d"})
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

type generator struct {
	ss []string
	i  int
	f  func(s []string, i int) string
}

func (seq *generator) Next() string {
	s := seq.f(seq.ss, seq.i)
	seq.i++
	return s
}
func sequenceFunc(ss ...string) (*generator, error) {
	if 0 == len(ss) {
		return nil, errors.New("sequence must have at least one element.")
	}
	return &generator{ss, 0, sequenceGen}, nil
}

func sequenceGen(ss []string, i int) string {
	if i >= len(ss) {
		return ss[len(ss)-1]
	}
	return ss[i]
}

func cycleFunc(ss ...string) (*generator, error) {
	if 0 == len(ss) {
		return nil, errors.New("cycle must have at least one element.")
	}
	return &generator{ss, 0, cycleGen}, nil
}

func cycleGen(ss []string, i int) string {
	return ss[i%len(ss)]
}
