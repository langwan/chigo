package main

import (
	"strings"
	"time"
)

func payload3() {
	for i := 0; i < 1000000; i++ {
		strings.Split("abc", "b")
	}
	time.Sleep(60 * time.Millisecond)
}
