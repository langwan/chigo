package main

import (
	"fmt"
	"github.com/langwan/langgo/helpers/gen"
)

func main() {
	str, err := helper_gen.RandString(32)
	if err != nil {
		return
	}
	fmt.Println(str)
}
