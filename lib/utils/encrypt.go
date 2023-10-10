package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

func generateSalt(n int) ([]byte, error) {
	b := make([]byte, n) 
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}
	return b, nil 
}

func EncryptAES256(data []byte, passphrase string) (encryptedData string, ivString string, err error) {
	salt, err := generateSalt(16)
	if err != nil {
		return "", "", err
	}
	
	key := pbkdf2.Key([]byte(passphrase), salt, 4096, 32, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", "", err
	}

	ciphertext := make([]byte, len(salt) + aes.BlockSize + len(data))
	copy(ciphertext[:len(salt)], salt)

	if len(ciphertext) < len(salt)+aes.BlockSize {
    return "", "", fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[len(salt):len(salt) + aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", "", err
	}
	copy(ciphertext[len(salt):len(salt)+aes.BlockSize], iv)

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[len(salt) + aes.BlockSize:], data)

  return base64.StdEncoding.EncodeToString(ciphertext), base64.StdEncoding.EncodeToString(iv), nil
}