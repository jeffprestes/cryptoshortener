package cryptonx

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
)

/*
Decrypter decripts text based on nonce and key
*/
func Decrypter(key string, nonce string, text string) (decryptedText string, err error) {
	err = nil
	textocifrado, _ := hex.DecodeString(text)
	nonceInterno, _ := hex.DecodeString(nonce)

	//fmt.Printf("[Decrypter] texto: [%v] \n", texto)
	//fmt.Printf("[Decrypter] textocifrado: [%v] \n", textocifrado)
	//fmt.Printf("[Decrypter] Nonce: [%v] \n", nonce)
	//fmt.Printf("[Decrypter] nonceInterno: [%v] \n", nonceInterno)

	aesgcm, err := GetAEAD(key)
	if err != nil {
		fmt.Printf("[Decrypter] Error during aesgcm generation: [%s]\n", err.Error())
		return
	}

	retorno, err := aesgcm.Open(nil, nonceInterno, textocifrado, nil)
	if err != nil {
		fmt.Printf("[Decrypter] Error during opening text operation: [%s]\n", err.Error())
		return
	}
	decryptedText = fmt.Sprintf("%s", retorno)
	return
}

//DecryptWithNonce it decrypts data and get the nonce from the encrypted data
func DecryptWithNonce(cipherText string, key string) (plainText string, err error) {
	err = nil
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Printf("[DecryptWithNonce] Error during Cipher generation. Error: %s\n", err.Error())
		return
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Printf("[DecryptWithNonce] Error during GCM generation. Error: %s\n", err.Error())
		return
	}

	tempCipherText, err := hex.DecodeString(cipherText)
	if err != nil {
		log.Printf("[DecryptWithNonce] Error during decoding cipherText string. Error: %s\n", err.Error())
		return
	}

	nonceSize := gcm.NonceSize()
	if len(tempCipherText) < nonceSize {
		err = errors.New("ciphertext too short")
		log.Printf("[DecryptWithNonce] Error checking nonceSize. Error: %s\n", err.Error())
		return
	}

	nonce, tempCipherText := tempCipherText[:nonceSize], tempCipherText[nonceSize:]
	bytesPlainText, err := gcm.Open(nil, nonce, tempCipherText, nil)
	if err != nil {
		log.Printf("[DecryptWithNonce] Error during opening encrypted text. Error: %s\n", err.Error())
		return
	}
	plainText = string(bytesPlainText)
	return
}
