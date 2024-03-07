package core

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

func New_Cracker(foundCH chan ResultPair, joinCH chan<- Signal, seedStr string) {

}

func (crack *Cracker) Start() {

}

func (crack *Cracker) ComputeHash(pt *string) string {
	var h hash.Hash

	switch crack.HashMethod {
	case SHA1:
		h = sha1.New()
	case SHA256:
		h = sha256.New()
	case SHA512:
		h = sha512.New()
	case MD5:
		h = md5.New()
	default:
		panic("Unknown Hash-Method")
	}

	h.Write([]byte(*pt))

	return hex.EncodeToString(h.Sum(nil))
}
