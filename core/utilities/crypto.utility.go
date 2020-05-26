package utilities

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
)

//import

func Hash(password string) ([]byte, error) {

	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

}

var initVector = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func EncryptString(key, text string) string {
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		panic(err)
	}

	plaintext := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, initVector)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func DecryptString(key, text string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	ciphertext, _ := base64.StdEncoding.DecodeString(text)
	cfb := cipher.NewCFBEncrypter(block, initVector)
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)
	return string(plaintext)
}
