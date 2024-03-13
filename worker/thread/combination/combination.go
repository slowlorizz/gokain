package combination

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"

	"github.com/slowlorizz/gokain/worker/thread/combination/charset"
	"github.com/slowlorizz/gokain/worker/thread/combination/wheel"
)

type (
	HashType uint8

	ComboWheels struct {
		Seeds wheel.Wheel
		Gears wheel.Wheel
	}

	Combination struct {
		Seed *wheel.Iterator
		Gear *Gear

		Source   charset.CharSet
		Wheels   ComboWheels
		HashType HashType
	}
)

const (
	SHA1   HashType = 1
	SHA256 HashType = 2
	SHA512 HashType = 3
	MD5             = 4
)

func New(seed []string, chrs charset.CharSet, ht HashType) *Combination {
	chrs.Build()

	cmb := Combination{Source: chrs, Wheels: ComboWheels{Seeds: *wheel.New(&seed), Gears: *wheel.New(&chrs.Chars)}, HashType: ht}
	cmb.Seed = cmb.Wheels.Seeds.NewIterator()
	cmb.Gear = &Gear{Itr: cmb.Wheels.Gears.NewIterator()}

	return &cmb
}

func (cmb *Combination) Next() string {
	str := cmb.Gear.Turn(cmb.Seed)
	return cmb.Seed.Item.Char + str
}

func (cmb *Combination) ComputeHash(str *string) string {
	var h hash.Hash

	switch cmb.HashType {
	case SHA1:
		h = sha1.New()
	case SHA256:
		h = sha256.New()
	case SHA512:
		h = sha512.New()
	case MD5:
		h = md5.New()
	default:
		panic("Unknown Hashtype")
	}

	h.Write([]byte(*str))

	return hex.EncodeToString(h.Sum(nil))
}

func (cmb *Combination) Cycle() (string, string) {
	str := cmb.Next()

	return str, cmb.ComputeHash(&str)
}