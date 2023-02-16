package main

import (
	"fmt"
	"math/big"
)

func main() {
	i := big.NewInt(100)
	fmt.Println("i", i)
	fmt.Println("i int64", i.Int64())

	m := big.NewInt(200)

	j := i.Add(i, m)

	fmt.Println("i", i, i)
	fmt.Println("j", j, j)

	fmt.Printf("i = %p, j = %p\n", i, j)

	i.Add(i, m)

	fmt.Println("i", i, i)
	fmt.Println("j", j, j)

	i = big.NewInt(100)
	mm := big.NewInt(0).Add(i, m)
	fmt.Println("mm", mm)
	fmt.Println("i", i)
	fmt.Printf("mm = %p, i = %p\n", mm, i)

	i = big.NewInt(0)
	i.SetString("1.99", 10)
	fmt.Println("i", i, i)

}
