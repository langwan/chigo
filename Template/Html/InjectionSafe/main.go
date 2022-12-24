package main

import (
	html_template "html/template"
	"os"
	text_template "text/template"
)

func main() {
	text, err := text_template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	if err != nil {
		panic(err)
	}
	err = text.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>\n")
	if err != nil {
		panic(err)
	}
	html, err := html_template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	if err != nil {
		panic(err)
	}
	err = html.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")
	if err != nil {
		panic(err)
	}
}
