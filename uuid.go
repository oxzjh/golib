package golib

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

func UUIDBytes() []byte {
	uuid := make([]byte, 16)
	io.ReadFull(rand.Reader, uuid)
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	return uuid
}

func UUID() string {
	return UUID2String(UUIDBytes())
}

func UUIDWithoutDash() string {
	buf := make([]byte, 32)
	hex.Encode(buf, UUIDBytes())
	return string(buf)
}

func UUID2String(uuid []byte) string {
	buf := make([]byte, 36)
	hex.Encode(buf[:8], uuid[:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], uuid[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], uuid[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], uuid[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], uuid[10:])
	return string(buf)
}
