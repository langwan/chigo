package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {

	tmpl, err := template.ParseFiles("./main.tmpl")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, nil)
	fmt.Println("")
	if err != nil {
		panic(err)
	}
}
