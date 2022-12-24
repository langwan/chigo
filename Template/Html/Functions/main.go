package main

import (
	"fmt"
	"os"
	"text/template"
)

func power(x int) int {
	return x * x
}

func main() {
	funcs := template.FuncMap{"power": power}
	glob, err := template.New("doc").Funcs(funcs).Parse(`{{ 2 | power}}`)
	if err != nil {
		panic(err)
	}
	err = glob.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("")
}
