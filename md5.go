package golib

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func Md5(val []byte) string {
	bytes := md5.Sum(val)
	return hex.EncodeToString(bytes[:])
}

func Md5Reader(r io.Reader) (string, error) {
	h := md5.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}
	bytes := h.Sum(nil)
	return hex.EncodeToString(bytes), nil
}

func Md5File(file string) (string, error) {
	r, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer r.Close()
	return Md5Reader(r)
}
