package determineifstringhalvesarealike

import "strings"

func halvesAreAlike(s string) bool {
	s = strings.ToLower(s)
	l := 0
	r := len(s) - 1
	count := 0
	for l < r {
		if s[l] == 'a' || s[l] == 'e' || s[l] == 'i' || s[l] == 'o' || s[l] == 'u' {
			count++
		}
		if s[r] == 'a' || s[r] == 'e' || s[r] == 'i' || s[r] == 'o' || s[r] == 'u' {
			count--
		}
		l++
		r--
	}

	return count == 0
}
