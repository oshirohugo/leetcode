package sortingthesentence

import "strings"

func sortSentence(s string) string {
	words := strings.Split(s, " ")

	sorted := make([]string, len(words), len(words))

	for _, word := range words {
		lastIndex := len(word) - 1
		i := word[lastIndex]
		sorted[i-'1'] = word[:lastIndex]

	}

	return strings.Join(sorted, " ")
}
