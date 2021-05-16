package rsa

import (
	"crypto/rand"
	"errors"
	"fmt"
)

func generatePrivateAndPublicKey() (*Key, *Key, error) {
	var e, p, q, d int64
	var err error

	p, q, err = choosePrimes()
	if err != nil {
		return nil, nil, err
	}

	n := p * q

	phi_n := eulersTotient(p, q)

	e, err = chooseSmallerCoprimeNumber(phi_n)
	if err != nil {
		return nil, nil, err
	}

	d, err = modularInverseSimple(e, phi_n)
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("p: ", p)
	fmt.Println("q: ", q)
	fmt.Println("n: ", n)
	fmt.Println("e: ", e)
	fmt.Println("d: ", d)
	fmt.Println("phi: ", phi_n)

	return &Key{Number: e, Mod: n}, &Key{Number: d, Mod: n}, nil

}

func choosePrimes() (int64, int64, error) {
	const msg = "p or q can not be represented as int64: try to choose smaller numbers"
	p, err1 := rand.Prime(rand.Reader, 6)
	q, err2 := rand.Prime(rand.Reader, 7)

	if err1 != nil {
		return 0, 0, err1
	}
	if err2 != nil {
		return 0, 0, err2
	}

	if !p.IsInt64() || !q.IsInt64() {
		return 0, 0, errors.New(msg)
	}

	return p.Int64(), q.Int64(), nil
}

func eulersTotient(p int64, q int64) int64 {
	return (p - 1) * (q - 1)
}

func chooseSmallerCoprimeNumber(phi_n int64) (int64, error) {
	test, err := rand.Prime(rand.Reader, 3)
	if err != nil {
		return 0, err
	}

	test_int64 := test.Int64()
	for phi_n%test_int64 == 0 || test_int64 >= phi_n {
		test, err = rand.Prime(rand.Reader, 3)
		test_int64 = test.Int64()
		if err != nil {
			return 0, err
		}
	}

	return test_int64, nil
}

// TODO: fix
func modularInverse(e, phi_n int64) (int64, error) {
	t := int64(0)
	r := phi_n
	newt := int64(1)
	newr := e

	for newr != 0 {
		quotient := r / newr

		prov_t := newt
		newt = t - quotient*prov_t
		t = prov_t

		prov_r := newr
		newr = r - quotient*prov_r
		r = prov_r
	}

	if r > 1 {
		return 0, errors.New("Not invertible")
	}

	if t < 0 {
		return int64(t) + phi_n, nil
	}

	return int64(t), nil
}

// modularInverseSimple computes the modular inverse of e via brute force
func modularInverseSimple(e, phi_n int64) (int64, error) {
	e %= phi_n
	for x := int64(1); x < phi_n; x++ {
		if (e*x)%phi_n == 1 {
			return x, nil
		}
	}

	return 0, errors.New("something went wrong")
}
