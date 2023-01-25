package truncatesentence

import "strings"

func truncateSentenceEasy(s string, k int) string {
	words := strings.Split(s, " ")
	words = words[:k]

	return strings.Join(words, " ")
}

func truncateSentence(s string, k int) string {
	for i, c := range s {
		if c == ' ' {
			k--
			if k == 0 {
				return s[:i]
			}
		}
	}

	return s
}
