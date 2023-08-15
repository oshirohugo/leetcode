package countthenumberofvowelstringsinrange

func vowelStrings(words []string, left int, right int) int {
	vowels := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}

	count := 0
	for i := left; i <= right; i++ {
		_, oks := vowels[words[i][0]]
		_, oke := vowels[words[i][len(words[i])-1]]

		if oks && oke {
			count++
		}
	}

	return count
}
