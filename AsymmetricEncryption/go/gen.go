package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

func GenKey(max int64, verbose bool) (*Key, error) {

	primes := GenerationOfPrimes(max)

	pi, err := rand.Int(rand.Reader, big.NewInt(int64(len(primes))))
	if err != nil {
		panic(err)
	}
	var qi *big.Int
	for {
		qi, err = rand.Int(rand.Reader, big.NewInt(int64(len(primes))))
		if err != nil {
			panic(err)
		}
		if pi.Cmp(qi) != 0 {
			break
		}
	}
	p := big.NewInt(primes[pi.Int64()])
	q := big.NewInt(primes[qi.Int64()])

	var d *big.Int

	n := big.NewInt(0).Mul(p, q)
	r := big.NewInt(0).Mul(big.NewInt(0).Sub(p, big.NewInt(1)), big.NewInt(0).Sub(q, big.NewInt(1)))

	e, d, ok, err := randFind(r, 999)

	//e = big.NewInt(19)
	//d = big.NewInt(7)

	if err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.New("unable to create key")
	}

	if verbose {
		fmt.Println("p = ", p)
		fmt.Println("q = ", q)
		fmt.Println("r = ", r)
		fmt.Println("-- secret --")
		fmt.Println("n = ", n)
		fmt.Println("e = ", e)
		fmt.Println("d = ", d)
	}

	return &Key{
		N: n,
		E: e,
		D: d,
	}, nil
}

func randFind(r *big.Int, tries int) (E, D *big.Int, ok bool, err error) {

	return _randFind(r, 0, 999)
}

func _randFind(r *big.Int, tries, max int) (E, D *big.Int, ok bool, err error) {
	if tries > max {
		return nil, nil, false, errors.New("trying too much")
	}

	Er, err := rand.Int(rand.Reader, big.NewInt(0).Sub(r, big.NewInt(2)))

	if err != nil {
		return nil, nil, false, err
	}

	E = big.NewInt(0).Add(Er, big.NewInt(2))

	if GCD(E, r).Cmp(big.NewInt(1)) != 0 {
		return _randFind(r, tries+1, max)
	}

	find := false
	var i int64

	for i = 2; i < r.Int64(); i++ {
		if big.NewInt(0).Mod(big.NewInt(0).Mul(big.NewInt(i), E), r).Cmp(big.NewInt(1)) == 0 && E.Cmp(big.NewInt(i)) != 0 {
			D = big.NewInt(i)
			find = true
			break
		}
	}

	if !find {
		return _randFind(r, tries+1, max)
	} else {
		return E, D, true, nil
	}

}
