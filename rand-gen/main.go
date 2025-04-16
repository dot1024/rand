package main

import (
	"flag"
	"fmt"

	"github.com/dot1024/rand"
)

const (
	typNum    = "n"
	typLower  = "l"
	typUpper  = "u"
	typLetter = "s"
	typAll    = "a"
)

const rndlen = 8

var (
	Version = "v1.0.1"
	BuiltAt = "2025-04-15T15:00:00+09:00"
)

var (
	rt string
	rl int
	vb bool
)

func init() {
	flag.StringVar(&rt, "rt", typAll, "type of random string, value contains n, l, u, s, a")
	flag.IntVar(&rl, "rl", rndlen, "length of random string")
	flag.BoolVar(&vb, "vb", false, "show all additional information")
	flag.Parse()
}

func main() {
	b := make([]byte, rl)
	switch rt {
	case typNum:
		rand.GenNum(b)
	case typLower:
		rand.GenLower(b)
	case typUpper:
		rand.GenUpper(b)
	case typLetter:
		rand.GenLetter(b)
	case typAll:
		rand.GenAll(b)
	default:
		fmt.Println(" → invalid arguments")
		flag.Usage()
	}
	if vb {
		fmt.Printf(" → len: %v\n", rl)
		fmt.Printf(" → str: %v\n", string(b))
	} else {
		println(" → str: " + string(b))
	}
}
