package main
import (
	"crypto/aes"
	"crypto/sha256"
	"crypto/cipher"
	"crypto/rand"
	"crypto/x509"
//	"encoding/hex"
	"fmt"
// "bufio"
	"io"
	"os"
	"github.com/urfave/cli/v2"
	"rohand2290/gopass/rsa_oaep"
	"rohand2290/gopass/error_handling"
)


func main() {
	/*
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
	*/
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "test",
				Aliases: []string{"t"},
				Usage: "test if app is working",
				Action: func(c *cli.Context) error {
					fmt.Println("test")
					return nil
				},
			},
			{
				Name: "add",
				Aliases: []string{"a"},
				Usage: "add a password",
				Action: func(c *cli.Context) error {
					fmt.Println("In the future this will add a password")
					return nil
				},
			},
			{
				Name: "init",
				Aliases: []string{"i"},
				Usage: "initializes the password manager",
				Action: func(c *cli.Context) error {
					passphrase := c.Args().Get(0)
					publicKey, privateKey := rsa_oaep.GetKeys() 
					fmt.Printf("Private Key: %x\n", x509.MarshalPKCS1PrivateKey(privateKey))
					fmt.Printf("Public Key: %x\n", x509.MarshalPKCS1PublicKey(publicKey))
					fmt.Printf("AES Key: %x\n", GetHash(passphrase))
					fmt.Printf("Encrypted Private Key: %x\n", EncryptWithString(passphrase, x509.MarshalPKCS1PrivateKey(privateKey)))
					return nil
				},
			},
			{
				Name: "get",
				Aliases: []string{"g"},
				Usage: "gets a password",
				Action: func (c* cli.Context) error {
					fmt.Println("This will get a password")
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	error_handling.CheckError(err)

}

func GetHash(s string) []byte {
	h := sha256.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

func EncryptWithString(s string, secret []byte) []byte {
	c, err := aes.NewCipher(GetHash(s))
	error_handling.CheckError(err)
	gcm, err := cipher.NewGCM(c)
	error_handling.CheckError(err)
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
  	fmt.Println(err)
  }
	return gcm.Seal(nonce, nonce, []byte(secret), nil)
}

func DecryptByteArr(encrypted, key []byte) []byte {
	c, err := aes.NewCipher(key) 
	error_handling.CheckError(err)
	gcm, err := cipher.NewGCM(c)
	error_handling.CheckError(err)
	nonceSize := gcm.NonceSize()
	nonce, encryptedText := encrypted[:nonceSize], encrypted[nonceSize:]
	decrypted, err := gcm.Open(nil, nonce, encryptedText, nil)
	error_handling.CheckError(nil)
	return decrypted
}


