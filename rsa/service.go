package rsa

import (
	"math/big"
)

type Service interface {
	EncryptMessage([]rune) []*big.Int
	DecryptMessage([]*big.Int) []rune
}

type service struct {
	key1 *Key
	key2 *Key
}

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
