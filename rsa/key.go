package rsa

// Key describes a private or public key, 2-dim-tuple
type Key struct {
	Number int64
	Mod    int64
}
