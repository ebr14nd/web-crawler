package wcserver

import (
	"net/url"
	"os"
	"path/filepath"
)

func readTestFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filepath.FromSlash("../../../testdata/" + filename))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func contains(s []*url.URL, str string) bool {
	for _, v := range s {
		if v.String() == str {
			return true
		}
	}
	return false
}
