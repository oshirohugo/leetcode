package splitstringsbyseparator

import "strings"

func splitWordsBySeparator(words []string, separator byte) []string {
	var ans []string

	for _, word := range words {
		s := strings.FieldsFunc(word, func(r rune) bool {
			return r == rune(separator)
		})
		ans = append(ans, s...)
	}

	return ans
}
