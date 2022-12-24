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
	foo := Foo{
		Name:  "chihuo",
		Value: 10,
	}
	tpl, err := template.New("doc").Parse("hello {{.Name}} {{if .Value}}{{.Value}} {{end}}\n")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(os.Stdout, foo)
	if err != nil {
		panic(err)
	}

	tpl, err = template.New("doc").Parse("hello {{.Name}} {{with .Value}}{{.}} {{end}}")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(os.Stdout, foo)
	if err != nil {
		panic(err)
	}
}
