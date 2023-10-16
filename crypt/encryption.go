package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"log"
)

const Key = "aVerySecureKeyyy"

type ImageCrypt struct {
	gcm cipher.AEAD
}

func New() ImageCrypt {
	block, err := aes.NewCipher([]byte(Key))
	if err != nil {
		log.Fatal(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}

	return ImageCrypt{gcm: gcm}
}

func (c ImageCrypt) Encrypt(plaintext []byte) ([]byte, error) {
	nonce := make([]byte, c.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := c.gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

func (c ImageCrypt) Decrypt(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < c.gcm.NonceSize() {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:c.gcm.NonceSize()], ciphertext[c.gcm.NonceSize():]
	plaintext, err := c.gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
