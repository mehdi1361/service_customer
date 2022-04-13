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
