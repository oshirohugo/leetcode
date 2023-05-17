package validpalindrome

import (
	"unicode"
)

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {

		if !isAlphanumeric(s[i]) {
			i++
			continue
		}

		if !isAlphanumeric(s[j]) {
			j--
			continue
		}

		if unicode.ToUpper(rune(s[i])) != unicode.ToUpper(rune(s[j])) {
			return false
		}
		i++
		j--
	}

	return true
}

func isAlphanumeric(c byte) bool {
	return c >= 'a' && c <= 'z' ||
		c >= 'A' && c <= 'Z' ||
		c >= '0' && c <= '9'
}
