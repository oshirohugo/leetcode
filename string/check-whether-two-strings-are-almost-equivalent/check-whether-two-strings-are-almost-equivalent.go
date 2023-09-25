package checkwhethertwostringsarealmostequivalent

func checkAlmostEquivalent(word1 string, word2 string) bool {
	count := [26]rune{}

	for _, c := range word1 {
		count[c-'a']++
	}
	for _, c := range word2 {
		count[c-'a']--
	}

	for _, el := range count {
		if el > 3 || el < -3 {
			return false
		}
	}

	return true
}
