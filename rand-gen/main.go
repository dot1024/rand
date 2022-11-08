package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/dot1024/rand"
)

const (
	defRandTyp = "1"
	defRandLen = 10
)

var (
	Version = "v0.0.0"
	BuiltAt = "2022-11-08T21:07:08+08:00"
)

var (
	rt string
	rl int
	vb bool
	st = time.Now()
)

func init() {
	flag.StringVar(&rt, "rt", defRandTyp, "type of random string, as binary number, base value is 1000, 0100, 0010, 0001")
	flag.IntVar(&rl, "rl", defRandLen, "length of random string")
	flag.BoolVar(&vb, "vb", false, "show all additional information")
	flag.Parse()
}

func main() {
	if vb {
		defer func(t time.Time) {
			fmt.Printf(" - spent: %.08fs\n", time.Since(t).Seconds())
		}(st)
	}

	if rt == "" || rl == 0 {
		fmt.Println(" - invalid arguments")
		flag.Usage()
		return
	}

	n, err := strconv.ParseInt(rt, 2, 0)
	if err != nil {
		fmt.Printf(" * invalid rt: %v\n", err)
		return
	}

	s, v, err := rand.Gen(int(n), rl)
	if err != nil {
		fmt.Printf(" * err: %v\n", err)
		return
	}

	if vb {
		fmt.Printf(" -  seed: %v\n", s)
		fmt.Printf(" -   len: %v\n", rl)
		fmt.Printf(" -   str: %v\n", v)
	} else {
		println(v)
	}
}
