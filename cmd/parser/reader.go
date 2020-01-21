package parser

import (
	"io/ioutil"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	file, fileErr := os.Open(path)
	if fileErr != nil {
		panic(fileErr)

	}

	return ioutil.ReadAll(file)
}
