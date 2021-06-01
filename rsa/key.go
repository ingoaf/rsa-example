package rsa

import "math/big"

// Key describes a private or public key, 2-dim-tuple
type Key struct {
	Number *big.Int
	Mod    *big.Int
}
