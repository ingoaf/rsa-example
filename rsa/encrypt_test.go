package rsa

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateKeys(t *testing.T) {
	key1, key2, err := generatePrivateAndPublicKey()
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, key1)
	assert.NotNil(t, key2)
	assert.Equal(t, key1.Mod.Cmp(key2.Mod), 0) // check mod equality
}

func TestKeysOnNumbers(t *testing.T) {
	var message, decryptedMessage int64
	tmp := big.NewInt(0)

	message = 5

	key1, key2, err := generatePrivateAndPublicKey()
	if err != nil {
		t.Fatal(err)
	}

	encrypted := tmp.Exp(big.NewInt(message), key1.Number, key1.Mod)
	decrypted := tmp.Exp(encrypted, key2.Number, key2.Mod)

	decryptedMessage = decrypted.Int64()

	assert.Equal(t, message, decryptedMessage)
}

func TestEulersTotient(t *testing.T) {

	prime1, prime2 := big.NewInt(11), big.NewInt(13)
	product := eulersTotient(prime1, prime2)

	assert.Equal(t, big.NewInt(120), product)
}
