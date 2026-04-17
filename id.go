package golib

import "math"

type ID struct {
	maxLen int
	midLen int
	shift  int
	max    uint64
	tail   uint64
	base   uint64
}

func (id *ID) Max() uint64 {
	return id.max
}

func (id *ID) Base() uint64 {
	return id.base
}

func (id *ID) SetBase(base uint64) {
	id.base = base
}

func (id *ID) Encrypt(n uint64) uint64 {
	m := n >> id.shift
	s := id.maxLen - 1
	for i := 0; i < id.shift; i++ {
		if t := n & (1 << i); t != 0 {
			m |= t << (s - i<<1)
		}
	}
	return id.base + m
}

func (id *ID) Decrypt(n uint64) uint64 {
	n -= id.base
	m := n & id.tail << id.shift
	s := id.midLen + 1 - id.shift
	for i := id.midLen; i < id.maxLen; i++ {
		if t := n & (1 << i); t != 0 {
			m |= t >> ((i-id.midLen)<<1 + s)
		}
	}
	return m
}

func NewID(n int) *ID {
	maxLen := n*3 + n/4
	midLen := maxLen / 2
	max := uint64(1 << maxLen)
	return &ID{maxLen, midLen, (maxLen + 1) / 2, max, 1<<midLen - 1, (uint64(math.Pow10(n)) - max) >> 1}
}
