package io

import (
	"crypto/sha256"
	"fmt"
	"io"
)

type RepeatByte byte

func (r RepeatByte) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = byte(r)
	}
	return len(p), nil
}

func Call_read() {
	testStream := RepeatByte('a')
	hasher := sha256.New()
	io.CopyN(hasher, testStream, 1000000)
	fmt.Printf("%x", hasher.Sum(nil))
}
