package utils

import (
	"math/rand"
	"strings"
)

func RandomCodeGenerate(length int) string {
	NumberOnly := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	result := make([]string, length)
	for i := 0; i < length; i++ {
		result[i] = NumberOnly[rand.Intn(len(NumberOnly))]
	}
	return strings.Join(result, "")
}

func HashMobile(phoneNumber string) string {
	result := make([]string, len(phoneNumber))
	for i, v := range phoneNumber {
		if i > 2 && i < 7 {
			result[i] = "*"
		} else {
			result[i] = string(v)
		}
	}
	return strings.Join(result, "")
}
