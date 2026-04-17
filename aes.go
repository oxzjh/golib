package golib

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func AesEncrypt(data, key, iv []byte) string {
	block, _ := aes.NewCipher(key)
	return base64.StdEncoding.EncodeToString(CBCDecrypt(block, data, iv))
}

func AesDecrypt(data string, key, iv []byte) ([]byte, error) {
	result, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	block, _ := aes.NewCipher(key)
	return CBCDecrypt(block, result, iv), nil
}

func CBCEncrypt(block cipher.Block, data, iv []byte) []byte {
	blockSize := block.BlockSize()
	paddingCount := blockSize - len(data)%blockSize
	paddingValue := byte(paddingCount)
	padding := make([]byte, paddingCount)
	for i := 0; i < paddingCount; i++ {
		padding[i] = paddingValue
	}
	data = append(data, padding...)
	result := make([]byte, len(data))
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(result, data)
	return result
}

func CBCDecrypt(block cipher.Block, data, iv []byte) []byte {
	result := make([]byte, len(data))
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(result, data)
	n := len(result)
	return result[:n-int(result[n-1])]
}
