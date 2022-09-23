package numberofdifferentintegersinastring

import (
	"regexp"
	"strings"
)

func numDifferentIntegers(word string) int {
	re := regexp.MustCompile(`[a-zA-Z]`)
	noLetters := re.ReplaceAllString(word, " ")

	nums := strings.Fields(noLetters)
	lookup := make(map[string]int)
	unique := 0

	for _, num := range nums {
		n := removeLeadingZeros(num)
		if _, ok := lookup[n]; !ok {
			unique++
		}
		lookup[n]++
	}

	return unique
}

func removeLeadingZeros(num string) string {
	re := regexp.MustCompile(`^0+`)
	return re.ReplaceAllString(num, "")
}
