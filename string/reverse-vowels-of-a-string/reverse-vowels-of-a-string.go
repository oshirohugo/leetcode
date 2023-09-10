package reversevowelsofastring

func reverseVowels(s string) string {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	ans := []byte(s)

	i := 0
	j := len(ans) - 1
	for i < j {
		for i < j && !vowels[ans[i]] {
			i++
		}
		for i < j && !vowels[ans[j]] {
			j--
		}

		aux := ans[i]
		ans[i] = ans[j]
		ans[j] = aux
		i++
		j--
	}

	return string(ans)
}
