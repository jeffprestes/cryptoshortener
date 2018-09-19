package cryptonx

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

/*
VerifySignature it verifies message signature
*/
func VerifySignature(appSecret string, bytes []byte, expectedSignature string) bool {
	mac := hmac.New(sha1.New, []byte(appSecret))
	mac.Write(bytes)
	if fmt.Sprintf("%x", mac.Sum(nil)) != expectedSignature {
		return false
	}
	return true
}
