package main

import "text/template"

func _panic(err error) {
	panic(err)
}

func main() {
	home, err := template.New("home").ParseFiles("./home.tmpl")
	_panic(err)
	page404, err := template.New("404").ParseFiles("./404.tmpl")
	_panic(err)

}
