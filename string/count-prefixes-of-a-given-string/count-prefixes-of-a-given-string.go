package countprefixesofagivenstring

import "strings"

func countPrefixes(words []string, s string) int {
	count := 0
	for _, word := range words {
		if strings.Index(s, word) == 0 {
			count++
		}
	}
	return count
}
