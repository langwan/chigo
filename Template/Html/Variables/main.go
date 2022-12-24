package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	tmpl, err := template.ParseFiles("./main.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, nil)
	fmt.Println("")
	if err != nil {
		panic(err)
	}
}
