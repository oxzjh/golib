package golib

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
)

func SignSHA256WithRSA(content string, key *rsa.PrivateKey) string {
	h := sha256.New()
	h.Write([]byte(content))
	hashed := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hashed)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(signature)
}

func VerifySignWithSHA256(content, sign string, key *rsa.PublicKey) error {
	h := sha256.New()
	h.Write([]byte(content))
	hashed := h.Sum(nil)
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(key, crypto.SHA256, hashed, signBytes)
}
