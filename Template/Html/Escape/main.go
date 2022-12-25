package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	const s = `"Fran & Freddie's Diner" <tasty@example.com>`
	v := []any{`"Fran & Freddie's Diner"`, ' ', `<tasty@example.com>`}

	fmt.Println(template.HTMLEscapeString(s))
	template.HTMLEscape(os.Stdout, []byte(s))
	fmt.Fprintln(os.Stdout, "")
	fmt.Println(template.HTMLEscaper(v...))

	fmt.Println(template.JSEscapeString(s))
	template.JSEscape(os.Stdout, []byte(s))
	fmt.Fprintln(os.Stdout, "")
	fmt.Println(template.JSEscaper(v...))

	fmt.Println(template.URLQueryEscaper(v...))
}
