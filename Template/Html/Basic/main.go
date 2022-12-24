package main

import (
	"os"
	"text/template"
)

type Data struct {
	Name string
}

func main() {

	data := Data{Name: "chihuo"}

	content, err := template.New("content").Parse("<h1>hello {{.Name}}</h1>")
	if err != nil {
		panic(err)
	}
	content.Execute(os.Stdout, data)
}
