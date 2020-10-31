package utils

import "io/ioutil"

// LoadJSONFixture loads a testing file
func LoadJSONFixture(filename string) (string, error) {
	filepath := filename

	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
