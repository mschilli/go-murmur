package murmur

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"errors"
)

const Version = "0.01"

const StoreFileName = ".murmur"

var StoreLocationFinder = findStorePath

func Lookup(name string) (string, error) {
    lookup, err := ReadStore()
    if err != nil {
	return "", err
    }

    pass, ok := lookup[name]
    if !ok {
	return "", errors.New(fmt.Sprintf("No entry for %s found", name))
    }

    return pass, nil
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

func ReadStore() (map[string]string, error) {
	data := make(map[string]string)

	path, err := StoreLocationFinder()
	if err != nil {
		return data, err
	}

	data, err = readJSONFile(path)

	return data, err
}
