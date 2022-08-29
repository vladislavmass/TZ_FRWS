package powwork

import (
	"encoding/base64"
	"fmt"
	"io"
	"strconv"

	"golang.org/x/crypto/sha3"
)

const maxIterations int = 1 << 23

func DoTheJob(task string, complexity int) (string, error) {
	counter := 0
	for {
		data := fmt.Sprintf("%s:%s", task, intToBase64(counter))
		hash, err := getSha3Hash(data)
		if err != nil {
			return "", err
		}
		if validateWork(hash, complexity) {
			return data, nil

		}
		if counter > maxIterations {
			return "", fmt.Errorf("max iterations reached")
		}
		counter++
	}

}
func getSha3Hash(s string) (string, error) {
	hash := sha3.New512()
	_, err := io.WriteString(hash, s)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
func validateWork(work string, complexity int) bool {
	for _, val := range work[:complexity] {
		if val != '0' {
			return false
		}
	}
	return true
}

func intToBase64(i int) string {
	return base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(i)))
}
