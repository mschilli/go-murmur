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

const Version = "0.02"

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
	lookup, err := m.readStore()
	if err != nil {
		return "", err
	}

	pass, ok := lookup[name]
	if !ok {
		return "", errors.New(fmt.Sprintf("No entry for %s found", name))
	}

	return pass, nil
}

func (m *MurmurStore) readStore() (map[string]string, error) {
	data := make(map[string]string)

	data, err := readJSONFile(m.FilePath)

	return data, err
}

func readJSONFile(path string) (map[string]string, error) {
	data := make(map[string]string)

	file, err := os.Open(path)
	if err != nil {
		return data, err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return data, err
	}

	err = yaml.Unmarshal(b, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
