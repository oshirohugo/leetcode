package checkifnumbersareascendinginasentence

import (
	"strconv"
	"strings"
)

func areNumbersAscending(s string) bool {
	tokens := strings.Fields(s)
	last := -1
	for _, t := range tokens {
		c, err := strconv.Atoi(t)
		if err != nil {
			continue
		}
		if c <= last {
			return false
		}
		last = c
	}
	return true
}
