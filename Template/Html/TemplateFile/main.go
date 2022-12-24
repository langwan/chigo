package main

import (
	"html/template"
	"os"
)

type Data struct {
	Name string
}

func main() {
	data := Data{Name: "chihuo"}
	tmpl, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
