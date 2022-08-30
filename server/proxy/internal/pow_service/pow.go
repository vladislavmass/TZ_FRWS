package powservice

import (
	"fmt"
	"strings"
)

type POWService struct {
	work       string
	complexity int
}

func CreatePOWService(complexity int) *POWService {
	return &POWService{
		work:       generateRandomString(40),
		complexity: complexity,
	}
}
func (pow POWService) GetWork() string {
	return fmt.Sprintf("%s:%d", pow.work, pow.complexity)
}

func (pow POWService) ValidateWork(workRezult string) bool {
	parsString := strings.Split(workRezult, ":")
	if pow.work != parsString[0] {
		return false
	}
	hash := getSha3Hash(workRezult)
	for _, val := range hash[:pow.complexity] {
		if val != '0' {
			return false
		}
	}
	return true

}
