package golib

func Mask(val, key []byte) []byte {
	keylen := len(key)
	result := make([]byte, len(val))
	for i, v := range val {
		result[i] = v ^ key[i%keylen]
	}
	return result
}

func MaskSelf(val, key []byte) {
	keylen := len(key)
	for i, v := range val {
		val[i] = v ^ key[i%keylen]
	}
}

func SimpleMask(val []byte, key byte) []byte {
	result := make([]byte, len(val))
	for i, v := range val {
		result[i] = v ^ key
	}
	return result
}

func SimpleMaskSelf(val []byte, key byte) {
	for i, v := range val {
		val[i] = v ^ key
	}
}

func ReverseMask(val []byte) []byte {
	result := make([]byte, len(val))
	for i, v := range val {
		result[i] = ^v
	}
	return result
}

func ReverseMaskSelf(val []byte) {
	for i, v := range val {
		val[i] = ^v
	}
}
