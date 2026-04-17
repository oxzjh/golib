package golib

import (
	"crypto/rand"
	"errors"
	"strings"
	"unicode/utf8"
)

const nonceChars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NonceString(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	l := byte(len(nonceChars))
	for i, v := range b {
		b[i] = nonceChars[v%l]
	}
	return string(b)
}

func NonceCode(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	for i, v := range b {
		b[i] = '0' + v%10
	}
	return string(b)
}

func Pad(s string, n int, r rune) string {
	l := utf8.RuneCountInString(s)
	if l < n {
		s = strings.Repeat(string(r), n-l) + s
	}
	return s
}

func Format62(i uint) string {
	b := make([]byte, 0, 16)
	for i >= 62 {
		b = append(b, nonceChars[i%62])
		i /= 62
	}
	return string(append(b, nonceChars[i]))
}

func Parse62(s string) (uint, error) {
	b, err := parseChar(s[0])
	if err != nil {
		return 0, err
	}
	r := uint(b)
	var t uint = 1
	for i := 1; i < len(s); i++ {
		b, err = parseChar(s[i])
		if err != nil {
			return 0, err
		}
		t *= 62
		r += t * uint(b)
	}
	return r, nil
}

func parseChar(b byte) (byte, error) {
	if '0' <= b && b <= '9' {
		return b - 48, nil
	}
	if 'a' <= b && b <= 'z' {
		return b - 87, nil
	}
	if 'A' <= b && b <= 'Z' {
		return b - 29, nil
	}
	return 0, errors.New("invalid syntax")
}
