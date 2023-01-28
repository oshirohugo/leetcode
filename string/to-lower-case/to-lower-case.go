package tolowercase

func toLowerCase(s string) string {
	lower := make([]rune, len(s))
	for i, c := range s {
		lower[i] = c
		if c >= 'A' && c <= 'Z' {
			lower[i] = ('A' - c + 'a')
		}

	}

	return string(lower)
}
