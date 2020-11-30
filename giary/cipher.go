package giary

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

type Client struct {
	gcm cipher.AEAD
}

func NewClient(password []byte) *Client {
	sha := sha256.New()
	sha.Write(password)
	key := sha.Sum(nil)

	c, err := aes.NewCipher(key)
	if err != nil {
		logger.Fatalln(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		logger.Fatalln(err)
	}
	return &Client{gcm}
}

func (c *Client) Seal(plaintext []byte) []byte {
	nonce := make([]byte, c.gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	return c.gcm.Seal(nonce, nonce, plaintext, nil)
}

func (c *Client) Open(ciphertext []byte) ([]byte, error) {
	nonceSize := c.gcm.NonceSize()
	nonce, body := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return c.gcm.Open(nil, nonce, body, nil)
}
