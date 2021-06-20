// Package crypto code for encrypting and decrypting is from https://play.golang.org/p/4FQBAeHgRs
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
)

var iv = []byte{12, 66, 41, 23, 90, 10, 67, 68, 05, 13, 90, 75, 99, 31, 56, 98}

//todo zaza include source reference

func Encrypt(text string) []byte {
	if text == "" {
		return nil
	}
	block, err := aes.NewCipher(signatureKey)
	if err != nil {
		log.Fatal("Error while using aes cipher to encrypt: ", err)
	}
	plaintext := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return ciphertext
}

func Decrypt(ciphertext []byte) string {
	if ciphertext == nil {
		return ""
	}
	block, err := aes.NewCipher(signatureKey)
	if err != nil {
		log.Fatal("Error while using aes cipher to decrypt: ", err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)
	return string(plaintext)
}
