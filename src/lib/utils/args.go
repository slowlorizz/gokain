package utils

import (
	"flag"
	"gokain/types"
)

type Args struct {
	Verbose  bool
	Hash     string
	Threads  int
	HashType types.HashType
}

func (a *Args) Load() {
	v := flag.Bool("v", false, "usage: -v | When defined verbose Output is shown (shows debug logs aswell)")
	h := flag.String("hash", "", "usage: --hash | the hash to crack")
	t := flag.Int("th", 1, "usage: -th | how many threads are created")
	ht := flag.String("type", "sha1", "the type of the hash to crack (LOWERCASE)\nSupports following:\nsha1, sha224, sha256, sha384, sha512\nmd4, md5\n| Default: sha1")

	flag.Parse()

	a.Verbose = *v
	a.Hash = *h
	a.Threads = *t

	switch *ht {
	case "sha1":
		a.HashType = types.SHA1
	case "sha256":
		a.HashType = types.SHA256
	case "sha512":
		a.HashType = types.SHA512
	case "md5":
		a.HashType = types.MD5

	default:
		panic("Unknown Hashtype!")
	}
}
