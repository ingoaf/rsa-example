package rsa

import (
	"crypto/rand"
	"math/big"
)

func generatePrivateAndPublicKey() (*Key, *Key, error) {
	var err error
	var p, q, e *big.Int
	n, d := big.NewInt(0), big.NewInt(0)

	p, q, err = choosePrimes()
	if err != nil {
		return nil, nil, err
	}

	n.Mul(p, q) // multiply primes

	phi_n := eulersTotient(p, q)

	e, err = chooseSmallerCoprimeNumber(phi_n)
	if err != nil {
		return nil, nil, err
	}

	d.ModInverse(e, phi_n)
	if err != nil {
		return nil, nil, err
	}

	return &Key{Number: e, Mod: n}, &Key{Number: d, Mod: n}, nil

}

func choosePrimes() (*big.Int, *big.Int, error) {
	const msg = "p or q can not be represented as int64: try to choose smaller numbers"
	p, err1 := rand.Prime(rand.Reader, 2000)
	q, err2 := rand.Prime(rand.Reader, 2000)

	if err1 != nil {
		return nil, nil, err1
	}
	if err2 != nil {
		return nil, nil, err2
	}

	return p, q, nil
}

func eulersTotient(p *big.Int, q *big.Int) *big.Int {
	var pMinusOne, qMinusOne, product big.Int

	one := big.NewInt(1)

	return product.Mul(pMinusOne.Sub(p, one), qMinusOne.Sub(q, one))
}

func chooseSmallerCoprimeNumber(phi_n *big.Int) (*big.Int, error) {

	test, err := rand.Prime(rand.Reader, 2000)
	if err != nil {
		return nil, err
	}

	return test, nil
}

func encryptRune(r rune, key *Key) *big.Int {
	message := big.NewInt(int64(r))
	return message.Exp(message, key.Number, key.Mod)
}

func decryptRune(encryptedRune *big.Int, key *Key) rune {
	encryptedRune = encryptedRune.Exp(encryptedRune, key.Number, key.Mod)

	return rune(encryptedRune.Int64())
}
