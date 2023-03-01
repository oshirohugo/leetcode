package reverseonlyletters

func reverseOnlyLetters(s string) string {
	i := 0
	j := len(s) - 1
	bts := []byte(s)

	for i < j {
		if !isLetter(bts[i]) {
			i++
			continue
		}
		if !isLetter(bts[j]) {
			j--
			continue
		}
		aux := s[i]
		bts[i] = s[j]
		bts[j] = aux
		i++
		j--
	}

	return string(bts)
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
