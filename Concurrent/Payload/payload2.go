package main

import "strings"

func payload2() {
	for i := 0; i < 2700000; i++ {
		strings.Split("abc", "b")
	}
}
