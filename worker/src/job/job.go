package job

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/slowlorizz/gokain/worker/src/thread/combination"
	"github.com/slowlorizz/gokain/worker/src/thread/combination/charset"
	"gopkg.in/yaml.v3"
)

type (
	Job struct {
		ID       int
		Threads  int
		Hash     string
		HashType combination.HashType
		Chars    []string
	}

	JobFile struct {
		Job struct {
			Id      int
			Threads int
			Hash    struct {
				String string
				Type   string
			}
			Charset map[string][]string
		}
	}
)

func New(path string) *Job {
	jf, err := ReadFile(path)

	if err != nil {
		panic(err.Error())
	}

	return &Job{ID: jf.Job.Id, Threads: jf.Job.Threads, Hash: jf.Job.Hash.String, HashType: GetHashType(jf.Job.Hash.Type), Chars: BuildCharset(jf.Job.Charset).Chars}
}

func ReadFile(path string) (*JobFile, error) {
	jf := &JobFile{}

	buf, err := os.ReadFile(GetPath(path))
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(buf, jf)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", path, err)
	}

	return jf, err
}

func GetPath(path string) string {
	p, _ := filepath.Abs(path)
	return p
}

func GetHashType(ht string) combination.HashType {
	switch ht {
	case "sha1", "SHA1":
		return combination.SHA1
	case "sha256", "SHA256":
		return combination.SHA256
	case "sha512", "SHA512":
		return combination.SHA512
	case "md5", "MD5":
		return combination.MD5
	default:
		panic("invalid Hash-Type in Job")
	}
}

func BuildCharset(selection map[string][]string) *charset.CharSet {
	// sobald Jobs von File Gelesen werden können hier ändern
	chrs, err := charset.New(charset.JobSelection{Charset: selection})

	if err != nil {
		panic(err.Error())
	}

	return chrs
}
