package murmur

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/user"
	"path"
)

const Version = "1.0.0"

type Murmur struct {
	FilePath string
}

const StoreFileName = ".murmur"

func NewMurmur() *Murmur {
	return &Murmur{}
}

func (m *Murmur) WithFilePath(path string) *Murmur {
	m.FilePath = path
	return m
}

func HomePath() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	p := path.Join(u.HomeDir, StoreFileName)

	return p, nil
}

func (m *Murmur) Lookup(name string) (string, error) {
	if len(m.FilePath) == 0 {
		path, err := HomePath()
		if err != nil {
			return "", err
		}
		m.FilePath = path
	}
	dict, err := readYAMLFile(m.FilePath)
	if err != nil {
		return "", err
	}

	pass, ok := dict[name]
	if !ok {
		return "", errors.New(fmt.Sprintf("No entry for %s found", name))
	}

	return pass, nil
}

func readYAMLFile(path string) (map[string]string, error) {
	data := make(map[string]string)

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return data, err
	}

	err = yaml.Unmarshal(raw, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
