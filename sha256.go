package golib

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func Sha256(val []byte) string {
	bytes := sha256.Sum256(val)
	return hex.EncodeToString(bytes[:])
}

func Sha256Reader(r io.Reader) (string, error) {
	h := sha256.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}
	bytes := h.Sum(nil)
	return hex.EncodeToString(bytes), nil
}

func Sha256File(file string) (string, error) {
	r, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer r.Close()
	return Sha256Reader(r)
}
