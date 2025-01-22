package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"golang.org/x/crypto/bcrypt"
)

// Factory function to create an EncryptionManager based on the strategy.
func New(key []byte) (*Encryption, error) {

	if len(key) == 0 {
		return nil, errors.New("encryption key can not be empty")
	}

	if len(key) == 16 || len(key) == 32 || len(key) == 64 {
		return &Encryption{key: key}, nil
	}

	return nil, errors.New("invalid encryption key length. Use 16, 32, or 64 bytes")
}

// Encryption implements EncryptionManager for AES encryption.
type Encryption struct {
	key []byte
}

func (e *Encryption) Encrypt(data string) (string, error) {
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return "", err
	}

	plainText := []byte(data)
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	// Encode cipherText to Base64 to ensure it is text-safe
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (e *Encryption) Decrypt(data string) (string, error) {
	// Decode Base64-encoded cipherText
	cipherText, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(e.key)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}

func (e *Encryption) Hash(data string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (e *Encryption) CompareHash(data, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(data))
	if err != nil {
		return false, err
	}

	return true, nil
}
