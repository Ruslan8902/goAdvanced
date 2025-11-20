package verify

import (
	"fmt"
	"os"
	"strings"
)

func ReadSomeFile(filename string) ([]byte, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	return file, nil
}

func IsJson(filename string) bool {
	return strings.HasSuffix(filename, ".json")
}
