package cryptonx

import "time"

/*
Nonces list of nonces to be used to decrypt
*/
var Nonces map[string]time.Time

func init() {
	Nonces = make(map[string]time.Time)
}

/*
IsNonceStillValid checks if the nonce is still valid
By default it valid for 10 minutes.
*/
func IsNonceStillValid(nonceKey string) bool {

	emitidoQuando := Nonces[nonceKey]

	if emitidoQuando.Nanosecond() < 10000 {
		return false
	}

	now := time.Now()

	if now.Sub(emitidoQuando).Minutes() > 10 {
		return false
	}

	return true
}
