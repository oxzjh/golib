package golib

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
)

func Sha1(val []byte) string {
	bytes := sha1.Sum(val)
	return hex.EncodeToString(bytes[:])
}

func Sha1Reader(r io.Reader) (string, error) {
	h := sha1.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}
	bytes := h.Sum(nil)
	return hex.EncodeToString(bytes), nil
}

func Sha1File(file string) (string, error) {
	r, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer r.Close()
	return Sha1Reader(r)
}
