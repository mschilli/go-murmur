package murmur

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

const Version = "1.0.0"

type MurmurStore struct {
	FilePath string
}

type MurmurOption func(*MurmurStore)

const StoreFileName = ".murmur"

func NewMurmur(opts ...MurmurOption) *MurmurStore {
	m := &MurmurStore{}

	for _, opt := range opts {
		opt(m)
	}

	if m.FilePath == "" {
		path, err := findStorePath()
		if err != nil {
			panic(err)
		}
		m.FilePath = path
	}

	return m
}

func WithFilePath(path string) MurmurOption {
	return func(m *MurmurStore) {
		m.FilePath = path
	}
}

func findStorePath() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	p := path.Join(u.HomeDir, StoreFileName)

	_, err = os.Stat(p)
	if err == nil {
		return p, nil
	}

	return "", err
}

func (m *MurmurStore) Lookup(name string) (string, error) {
	dict, err := readJSONFile(m.FilePath)
	if err != nil {
		return "", err
	}

	pass, ok := dict[name]
	if !ok {
		return "", errors.New(fmt.Sprintf("No entry for %s found", name))
	}

	return pass, nil
}

func readJSONFile(path string) (map[string]string, error) {
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
