package powservice

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/crypto/sha3"
)

func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"~=+%^*/()[]{}/!@#$?|" +
		"0123456789"
	var work strings.Builder
	for i := 0; i < length; i++ {
		work.WriteByte(all[rand.Intn(len(all))])
	}
	return work.String()
}

func getSha3Hash(s string) string {
	hash := sha3.New512()
	_, err := io.WriteString(hash, s)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}
