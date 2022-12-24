package main

import (
	"fmt"
	"os"
	"text/template"
)

func _panic(err error) {
	if err != nil {
		panic(err)
	}
}

type DataHome struct {
	Name string
}

func main() {
	var err error
	defer func() {
		if err != nil {
			panic(err)
		}
	}()
	tmpl, err := template.ParseFiles("./layout.tmpl", "./home.tmpl")
	tmpl.Execute(os.Stdout, &DataHome{Name: "chihuo"})
	fmt.Println("")
	fmt.Println("")

	tmpl, err = template.ParseFiles("./layout.tmpl", "./404.tmpl")
	tmpl.Execute(os.Stdout, &DataHome{Name: "chihuo"})
	fmt.Println("")
}
