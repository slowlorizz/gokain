package charset

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var CONFIG_FILE string = GetPath("./config/charset.config.yml") //"../../../../config/charset.config.yml"

type CharSet struct {
	ConfigFile string
	Config     Config
	Selection  map[string][]string
	Chars      []string
}

type Config struct {
	Charset map[string]map[string][]string
}

type JobSelection struct {
	Charset map[string][]string
}

func New(js JobSelection) (*CharSet, error) {
	chrs := CharSet{
		ConfigFile: CONFIG_FILE,
		Config:     Config{},
		Selection:  js.Charset,
		Chars:      make([]string, 0),
	}

	err := chrs.Config.ReadFile(chrs.ConfigFile)
	chrs.Include(chrs.Selection)

	return &chrs, err
}

func (C *Config) ReadFile(path string) error {
	buf, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, C)
	if err != nil {
		return fmt.Errorf("in file %q: %w", path, err)
	}

	return err
}

func (chrs *CharSet) Include(s map[string][]string) {
	for k, v := range s {
		for _, x := range v {
			chrs.Chars = append(chrs.Chars, chrs.Config.Charset[k][x]...)
		}
	}
}

func GetPath(path string) string {
	p, _ := filepath.Abs(path)
	return p
}
