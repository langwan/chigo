package main

import (
	"os"
	"text/template"
)

type Data struct {
	Title   string
	Content string
}

func main() {

	layout, err := template.ParseFiles("./tpl/layout.html")
	if err != nil {
		panic(err)
	}

	layout, err = layout.ParseFiles("./tpl/home.html")
	if err != nil {
		panic(err)
	}

	err = layout.Execute(os.Stdout, Data{Title: "title", Content: "chihuo"})
	if err != nil {
		panic(err)
	}
}
