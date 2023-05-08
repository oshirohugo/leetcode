package firstlettertoappeartwice

func repeatedCharacter(s string) byte {
	lookup := make(map[rune]bool)
	for i, c := range s {
		if lookup[c] {
			return s[i]
		} else {
			lookup[c] = true
		}
	}

	return s[0]
}
