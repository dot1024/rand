package rand

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	TypNum     = 1 << 0 // TypNum source type for numeric characters
	TypLower   = 1 << 1 // TypLower source type for lowercase letters
	TypUpper   = 1 << 2 // TypUpper source type for uppercase letters
	TypSpecial = 1 << 3 // TypSpecial source type for special characters (keyboard visible characters)
	TypAll     = 1 << 4 // TypAll source type for all of aboving
)

const (
	minRandLen = 4      // minRandLen min length for generator
	maxRandLen = 1 << 8 // maxRandLen max length for generator
)

// mapping for source characters
var ssMap = map[int]string{
	TypNum:     "0123456789",                         // numeric characters
	TypLower:   "abcdefghijklmnopqrstuvwxyz",         // lowercase letters
	TypUpper:   "ABCDEFGHIJKLMNOPQRSTUVWXYZ",         // uppercase letters
	TypSpecial: "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~", // special characters
}

// Gen generate string
//
//   - [in] t: random source type, the binary number is option, eg: 1000, 0100, 0010, 0001, 1100 and so on
//   - [in] l: random string length, the number should belong to closed interval [4, 256] <- [minRandLen, maxRandLen]
//   - [out] ss: source string
//   - [out] sv: random string be generated
//   - [out] err: errors that occurred during program processing
func Gen(t, l int) (ss string, sv string, err error) {
	if l < minRandLen || l > maxRandLen {
		err = fmt.Errorf("random length must be belong to [%d, %d]", minRandLen, maxRandLen)
		return
	}

	if t > TypAll {
		err = errors.New("invalid typ")
		return
	}

	if t&TypNum == TypNum {
		ss += ssMap[TypNum]
	}
	if t&TypLower == TypLower {
		ss += ssMap[TypLower]
	}
	if t&TypUpper == TypUpper {
		ss += ssMap[TypUpper]
	}
	if t&TypSpecial == TypSpecial {
		ss += ssMap[TypSpecial]
	}

	if ss == "" {
		err = fmt.Errorf("invalid typ: %v", t)
		return
	}

	sl := len(ss)

	sb := strings.Builder{}
	sb.Grow(l)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < l; i++ {
		sb.WriteByte(ss[rand.Intn(sl)])
	}

	sv = sb.String()

	return
}
