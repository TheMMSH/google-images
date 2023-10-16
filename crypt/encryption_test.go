package crypt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImageCryptCanDecryptWhatItEncrypts(t *testing.T) {
	asserts := assert.New(t)

	sut := New()

	plainText := "My Name Is Mohammad Mehdi!"

	cypherText, _ := sut.Encrypt([]byte(plainText))
	decryptedText, _ := sut.Decrypt(cypherText)

	asserts.Equal(plainText, string(decryptedText))
}
