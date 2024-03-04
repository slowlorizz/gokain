package main

import (
	"flag"
)

type Args struct {
	verbose bool
	hash    string
}

func (a *Args) Load() {
	v := flag.Bool("v", false, "usage: -v | When defined verbose Output is shown (shows debug logs aswell)")
	h := flag.String("hash", "", "usage: --hash | the hash to crack")

	flag.Parse()

	a.verbose = *v
	a.hash = *h
}
