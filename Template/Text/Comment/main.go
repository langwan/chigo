package main

import (
	"os"
	"text/template"
)

type Foo struct {
	Name  string
	Value int
}

func main() {
	tpl, err := template.New("doc").Parse("my name is {{/* a comment */}}chihuo.")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}
