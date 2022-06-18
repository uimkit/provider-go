package uim

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func checkSignature(result, secret string, body []byte) bool {
	expected := hmacBytes(body, []byte(secret))
	resultBytes, err := hex.DecodeString(result)
	if err != nil {
		return false
	}
	return hmac.Equal(expected, resultBytes)
}

func hmacBytes(toSign, secret []byte) []byte {
	_authSignature := hmac.New(sha256.New, secret)
	_authSignature.Write(toSign)
	return _authSignature.Sum(nil)
}
