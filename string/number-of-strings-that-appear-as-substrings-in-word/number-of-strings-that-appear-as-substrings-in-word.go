package numberofstringsthatappearassubstringsinword

import (
	"strings"
)

func numOfStrings(patterns []string, word string) int {
	counter := 0
	for _, p := range patterns {
		if strings.Contains(word, p) {
			counter++
		}
	}
	return counter
}
