package rsa

import (
	"crypto/rand"
	"math/big"
)

// Service represents an encryption technique with encryption/decryption capability
type Service interface {
	EncryptMessage([]rune) []*big.Int
	DecryptMessage([]*big.Int) []rune
}

type service struct {
	key1 *Key
	key2 *Key
}

// NewService generates an new rsa instance with keys
func NewService() Service {
	key1, key2, err := generatePrivateAndPublicKey()
	if err != nil {
		panic("can not generate private and public key: " + err.Error())
	}
	return &service{key1: key1, key2: key2}
}

func (s *service) EncryptMessage(message []rune) []*big.Int {

	var encryptedMessage []*big.Int

	for _, runeElement := range message {

		encryptedMessage = append(encryptedMessage, encryptRune(runeElement, s.key1))
	}

	return encryptedMessage
}

func (s *service) DecryptMessage(encryptedMessage []*big.Int) []rune {
	var decryptedMessage []rune

	for _, encryptedRune := range encryptedMessage {
		decryptedMessage = append(decryptedMessage, decryptRune(encryptedRune, s.key2))
	}

	return decryptedMessage
}

func generatePrivateAndPublicKey() (*Key, *Key, error) {
	var err error
	var p, q, e *big.Int
	n, d := big.NewInt(0), big.NewInt(0)

	p, q, err = choosePrimes()
	if err != nil {
		return nil, nil, err
	}

	n.Mul(p, q) // multiply primes

	phiN := eulersTotient(p, q)

	e, err = chooseSmallerCoprimeNumber(phiN)
	if err != nil {
		return nil, nil, err
	}

	d.ModInverse(e, phiN)
	if err != nil {
		return nil, nil, err
	}

	return &Key{Number: e, Mod: n}, &Key{Number: d, Mod: n}, nil

}

func choosePrimes() (*big.Int, *big.Int, error) {
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

func chooseSmallerCoprimeNumber(phiN *big.Int) (*big.Int, error) {

	e, err := rand.Prime(rand.Reader, 2000)
	if err != nil {
		return nil, err
	}

	if e.Cmp(phiN) == -1 {
		return e, nil
	}

	e, err = chooseSmallerCoprimeNumber(phiN)
	if err != nil {
		return e, nil
	}

	return e, nil
}

func encryptRune(r rune, key *Key) *big.Int {
	message := big.NewInt(int64(r))
	return message.Exp(message, key.Number, key.Mod)
}

func decryptRune(encryptedRune *big.Int, key *Key) rune {
	encryptedRune = encryptedRune.Exp(encryptedRune, key.Number, key.Mod)

	return rune(encryptedRune.Int64())
}
