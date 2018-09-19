package cryptonx

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

/*
Encrypter Encrypts text based on AES Algorythim and return the encrypted text and his nonce
*/
func Encrypter(key string, text string) (encryptedText string, nonce string, err error) {
	err = nil
	nonceTemp := make([]byte, 12)
	if _, err = io.ReadFull(rand.Reader, nonceTemp); err != nil {
		fmt.Printf("[Encrypter] Error during nonce generation: [%s]\n", err.Error())
		return
	}

	aesgcm, err := GetAEAD(key)
	if err != nil {
		fmt.Printf("[Encrypter] Error during AEAD generation: [%s]\n", err.Error())
		return
	}
	encryptedText = fmt.Sprintf("%x", aesgcm.Seal(nil, nonceTemp, []byte(text), nil))
	nonce = fmt.Sprintf("%x", nonceTemp)
	return
}

//EncryptWithNonce it encrypts data and embbeded the nonce within the encrypted data
func EncryptWithNonce(plainText string, key string) (cypherText string, err error) {
	err = nil
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Printf("[EncryptWithNonce] Error during Cypher generation. Error: %s\n", err.Error())
		return
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Printf("[EncryptWithNonce] Error during GCM generation. Error: %s\n", err.Error())
		return
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Printf("[EncryptWithNonce] Error during nonce reading. Error: %s\n", err.Error())
		return
	}

	cypherText = hex.EncodeToString(gcm.Seal(nonce, nonce, []byte(plainText), nil))
	return
}
