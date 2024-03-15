package args

import (
	"flag"

	"github.com/slowlorizz/gokain/worker/src/exc"
	"github.com/slowlorizz/gokain/worker/src/thread/combination"
)

var Threads int
var Hash string
var HashType combination.HashType

func Load() *exc.Exception {
	t := flag.Int("th", 1, "usage: -th | how many threads are created")
	h := flag.String("hash", "", "usage: --hash | the hash to crack")
	ht := flag.String("type", "sha1", "the type of the hash to crack (LOWERCASE)\nSupports following:\nsha1, sha224, sha256, sha384, sha512\nmd4, md5\n| Default: sha1")

	flag.Parse()

	Hash = *h
	Threads = *t

	switch *ht {
	case "sha1", "SHA1":
		HashType = combination.SHA1
	case "sha256", "SHA256":
		HashType = combination.SHA256
	case "sha512", "SHA512":
		HashType = combination.SHA512
	case "md5", "MD5":
		HashType = combination.MD5

	default:
		return exc.New(exc.UserArgsException, "unsupported_hash_type", "Unsupported Hashtype in Arguments", exc.Data{
			"hash-type": *ht,
		})
	}

	if len(Hash) < 1 || Hash == "" {
		return exc.New(exc.UserArgsException, "invalid_hash", "Invalid Hash-Argument Value", exc.Data{
			"hash": Hash,
		})
	}

	if Threads < 1 {
		return exc.New(exc.UserArgsException, "insufficient_thread_count", "App needs atleast 1 Thread", exc.Data{
			"thread-count": Threads,
		})
	}

	return nil
}
