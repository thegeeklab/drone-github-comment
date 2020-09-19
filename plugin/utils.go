package plugin

import (
	"io/ioutil"
	"os"
)

func readStringOrFile(input string) (string, error) {
	if len(input) > 255 {
		return input, nil
	}
	// Check if input is a file path
	if _, err := os.Stat(input); err != nil && os.IsNotExist(err) {
		// No file found => use input as result
		return input, nil
	} else if err != nil {
		return "", err
	}
	result, err := ioutil.ReadFile(input)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
