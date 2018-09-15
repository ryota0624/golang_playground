package common

import (
	"io/ioutil"

	"github.com/rakyll/statik/fs"
)

func LoadStatikFS(path string) (string, error) {

	statikFS, err := fs.New()

	if err != nil {
		return "", err
	}

	file, err := statikFS.Open(path)

	if err != nil {
		return "", err
	}

	queryAsByte, err := ioutil.ReadAll(file)

	if err != nil {
		return "", err
	}

	return string(queryAsByte), nil
}
