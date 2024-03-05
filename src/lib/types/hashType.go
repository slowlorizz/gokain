package types

type HashType uint

const (
	SHA1   HashType = 0
	SHA256 HashType = 1
	SHA512 HashType = 2
	MD5    HashType = 3
)

func (ht HashType) String() string {
	switch ht {
	case SHA1:
		return "sha1"
	case SHA256:
		return "sha256"
	case SHA512:
		return "sha512"
	case MD5:
		return "md5"
	default:
		return "unknown"
	}
}
