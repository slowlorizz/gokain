package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"gokain/types"
	"hash"
)

func ComputeHash(txt string, ht types.HashType) string {
	var h hash.Hash

	switch ht {
	case types.SHA1:
		h = sha1.New()
	case types.SHA256:
		h = sha256.New()
	case types.SHA512:
		h = sha512.New()
	case types.MD5:
		h = md5.New()
	default:
		panic("Unknown Hashtype")
	}

	h.Write([]byte(txt))

	return hex.EncodeToString(h.Sum(nil))
}
