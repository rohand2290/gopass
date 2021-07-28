package main
import (
	"crypto/aes"
	"crypto/sha256"
	"crypto/cipher"
	"crypto/rand"
//	"encoding/hex"
	"fmt"
	"bufio"
	"io"
	"os"
)
func main() {
	passphrase := ""
	secret := ""
	fmt.Println("Enter passphrase: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		passphrase = scanner.Text()
	}
	fmt.Println("Enter secret: ")
	if scanner.Scan() {
		secret = scanner.Text()
	}

	encryptedSecret := EncryptWithString(passphrase, secret)
	decryptedSecret := DecryptByteArr(encryptedSecret, GetHash(passphrase))
	fmt.Printf("Encrypted String: %x\n", encryptedSecret)
	fmt.Printf("Decrypted String: %s\n", decryptedSecret)

}

func GetHash(s string) []byte {
	h := sha256.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}
func EncryptWithString(s, secret string) []byte {
	c, err := aes.NewCipher(GetHash(s))
	CheckError(err)
	gcm, err := cipher.NewGCM(c)
	CheckError(err)
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
  	fmt.Println(err)
  }
	return gcm.Seal(nonce, nonce, []byte(secret), nil)
}
func DecryptByteArr(encrypted, key []byte) []byte {
	c, err := aes.NewCipher(key) 
	CheckError(err)
	gcm, err := cipher.NewGCM(c)
	CheckError(err)
	nonceSize := gcm.NonceSize()
	nonce, encryptedText := encrypted[:nonceSize], encrypted[nonceSize:]
	decrypted, err := gcm.Open(nil, nonce, encryptedText, nil)
	CheckError(nil)
	return decrypted
	
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
