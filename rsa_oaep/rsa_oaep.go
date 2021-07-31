package rsa_oaep

import (
	"crypto/rand"
	"crypto/rsa"
	"rohand2290/gopass/error_handling"
)
func GetKeys() (*rsa.PublicKey, *rsa.PrivateKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 3072)
	error_handling.CheckError(err)
	publicKey := &privateKey.PublicKey
	return publicKey, privateKey

}
