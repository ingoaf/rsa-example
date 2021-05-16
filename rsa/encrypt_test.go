package rsa

import (
	"fmt"
	"testing"
)

func TestGenerateKeys(t *testing.T) {
	fmt.Println(generatePrivateAndPublicKey())
}
