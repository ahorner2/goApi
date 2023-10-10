package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

func DecryptAES256(encodedData string, passphrase string) (decryptedData string, err error) {
	// decode the base64
	data, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return "", err
	}

	// Check minimum length (salt + block size)
	if len(data) < 16+aes.BlockSize {
		return "", errors.New("encoded data too short")
	}

	// split salt, iv and ciphertext
	salt := data[:16]
	iv := data[16 : 16+aes.BlockSize]
	ciphertext := data[16+aes.BlockSize:]

	// derive key using the passphrase and salt
	key := pbkdf2.Key([]byte(passphrase), salt, 4096, 32, sha256.New)

	// create the AES cipher with the derived key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Check ciphertext length
	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	// decrypt the data using CFB mode
	stream := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}