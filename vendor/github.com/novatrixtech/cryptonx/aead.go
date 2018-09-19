package cryptonx

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

/*
GetAEAD gets the authenticated encryption object
*/
func GetAEAD(key string) (cipher.AEAD, error) {
	bloco, erro := aes.NewCipher([]byte(key))
	if erro != nil {
		fmt.Printf("[GetAEAD] Erro ao gerar o bloco: [%s]\n", erro.Error())
		return nil, erro
	}

	aesgcm, erro := cipher.NewGCM(bloco)
	if erro != nil {
		fmt.Printf("[GetAEAD] Erro ao gerar o aesgcm: [%s]\n", erro.Error())
		return nil, erro
	}

	return aesgcm, nil
}
