# cryptonx
Crypto is a lib to encrypt and decrypt messages using Go Language. It's very simple and straightforward.

## Example
```go
package main

import (
	"fmt"
	"strconv"

	"github.com/novatrixtech/cryptonx"
)

func main() {
	var key = "12345678901234567890123456789012"
	var teste int = 342
	//var text = "342"
	textEncrypted, nonce, err := cryptonx.Encrypter(key, strconv.Itoa(teste))
	if err != nil {
		panic(err)
	}
	fmt.Printf("=====\n")
	fmt.Println("Encrypted text is: ", textEncrypted)
	fmt.Println("The nonce (SALT) is: ", nonce)
	textDecrypted, err := cryptonx.Decrypter(key, nonce, textEncrypted)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypted text is: ", textDecrypted)

	//Encrypt and Decrypt text with the nonce (SALT) embedded at encrypted text
	fmt.Printf("=====\n")
	textEncrypted, err = cryptonx.EncryptWithNonce(strconv.Itoa(teste), key)
	if err != nil {
		panic(err)
	}
	fmt.Println("Encrypted text with nonce is: ", textEncrypted)
	textDecrypted, err = cryptonx.DecryptWithNonce(textEncrypted, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypted text with nonce is: ", textDecrypted)
	fmt.Printf("=====\n")
}
```

