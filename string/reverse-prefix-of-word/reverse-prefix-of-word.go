package reverseprefixofword

func reversePrefix(word string, ch byte) string {

	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			return reverse([]byte(word), i)
		}
	}

	return word
}

func reverse(word []byte, i int) string {
	j := 0
	for j < i {
		aux := word[j]
		word[j] = word[i]
		word[i] = aux
		j++
		i--
	}

	return string(word)
}
