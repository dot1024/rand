package rand

import (
	randv2 "math/rand/v2"
)

const (
	nsrc = "0123456789"                         // numeric characters
	lsrc = "abcdefghijklmnopqrstuvwxyz"         // lowercase letters
	usrc = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"         // uppercase letters
	ssrc = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~" // special characters
	asrc = nsrc + lsrc + usrc + ssrc            // all characters
)

// src source
type src string

// gen generate random data
func (s src) gen(b []byte) {
	if len(s) == 0 {
		return
	}
	if len(b) == 0 {
		return
	}
	for i := range b {
		b[i] = s[randv2.IntN(len(s))]
	}
}

// GenLetter generate random data using number strings
func GenNum(b []byte) {
	src(nsrc).gen(b)
}

// GenLetter generate random data with lower letters
func GenLower(b []byte) {
	src(lsrc).gen(b)
}

// GenLetter generate random data with upper letters
func GenUpper(b []byte) {
	src(usrc).gen(b)
}

// GenLetter generate random data using letters
func GenLetter(b []byte) {
	src(lsrc + usrc).gen(b)
}

// GenAll generate random data using all of string
func GenAll(b []byte) {
	src(asrc).gen(b)
}
