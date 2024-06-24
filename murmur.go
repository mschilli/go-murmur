package murmur

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/user"
	"path"
)

const Version = "1.0.0"

// Read secrets from a .murmur YAML file
type Murmur struct {
	FilePath string
}

const StoreFileName = ".murmur"

// Create a new instance
func NewMurmur() *Murmur {
	return &Murmur{}
}

// Set the .murmur file path manually
func (m *Murmur) WithFilePath(path string) *Murmur {
	m.FilePath = path
	return m
}

func homePath() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	p := path.Join(u.HomeDir, StoreFileName)

	return p, nil
}

// Look up a .murmur key by name and return its value
func (m *Murmur) Lookup(name string) (string, error) {
	if len(m.FilePath) == 0 {
		path, err := homePath()
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
		return "", fmt.Errorf("No entry found for %s", name)
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
