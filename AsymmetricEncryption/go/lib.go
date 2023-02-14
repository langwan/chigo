package main

import "math/big"

func GCD(a, b *big.Int) *big.Int {
	for b.Cmp(big.NewInt(0)) != 0 {
		t := b
		b = big.NewInt(0).Mod(a, b)
		a = t
	}
	return a
}

func Encrypt(key *Key, data *big.Int) *big.Int {
	return big.NewInt(0).Exp(data, key.E, key.N)
}

func Decrypt(key *Key, data *big.Int) *big.Int {
	return big.NewInt(0).Exp(data, key.D, key.N)
}

// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
func GenerationOfPrimes(max int64) []int64 {

	a := make([]bool, max+1)
	var i int64
	for i = 2; i < max; i++ {
		a[i] = true
	}

	for i = 2; i < max; i++ {
		if a[i] {
			for j := i; j < (max/i)+1; j++ {
				a[i*j] = false

			}
		}
	}
	total := 0

	for i = 2; i < max; i++ {
		if a[i] {
			total++
		}
	}

	output := make([]int64, total)
	index := 0
	for i = 2; i < max; i++ {
		if a[i] {
			output[index] = i
			index++
		}
	}

	return output
}

type Key struct {
	N *big.Int
	E *big.Int
	D *big.Int
}
