package main

import (
	"flag"
)

type Args struct {
	verbose bool
}

func (a *Args) Load() {
	v := flag.Bool("v", false, "usage: -v | When defined verbose Output is shown (shows debug logs aswell)")

	flag.Parse()

	a.verbose = *v
}
