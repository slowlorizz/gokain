package main

import (
	"flag"

	"github.com/slowlorizz/gokain/worker/thread"
	"github.com/slowlorizz/gokain/worker/thread/combination"
	"github.com/slowlorizz/gokain/worker/thread/combination/charset"
	"github.com/slowlorizz/gokain/worker/ui"
)

var Result thread.ResultPair

func main() {
	args := Args{}
	args.Load()

	ui.Init()

	chrs := charset.New(true, true, true, true, false, false, false)

	seeds := make([][]string, args.Threads)

	for i := 0; i < args.Threads; i++ {
		seeds[i] = make([]string, 0)
	}

	for i, c := range chrs.Chars {
		seeds[i%args.Threads] = append(seeds[i%args.Threads], c)
	}

	for i, v := range seeds {
		th := thread.New(i+1, args.Hash, v, args.HashType, *chrs, thread.FOUND_CH, thread.JOIN_CH)
		go th.Start()
	}

	for {
		select {
		case e := <-ui.HANDLER.Events:
			if ui.HANDLER.HandleEvent(&e) {
				thread.StopAll()
				return
			}
		case res := <-thread.FOUND_CH:
			Result = res
			thread.StopAll()
			ui.Render()
		default:
			ui.Render()
		}
	}
}

type (
	Args struct {
		Threads  int
		Hash     string
		HashType combination.HashType
	}
)

func (args *Args) Load() {
	t := flag.Int("th", 1, "usage: -th | how many threads are created")
	h := flag.String("hash", "", "usage: --hash | the hash to crack")
	ht := flag.String("type", "sha1", "the type of the hash to crack (LOWERCASE)\nSupports following:\nsha1, sha224, sha256, sha384, sha512\nmd4, md5\n| Default: sha1")

	flag.Parse()

	args.Hash = *h
	args.Threads = *t

	switch *ht {
	case "sha1", "SHA1":
		args.HashType = combination.SHA1
	case "sha256", "SHA256":
		args.HashType = combination.SHA256
	case "sha512", "SHA512":
		args.HashType = combination.SHA512
	case "md5", "MD5":
		args.HashType = combination.MD5

	default:
		panic("Unknown Hashtype!")
	}
}
