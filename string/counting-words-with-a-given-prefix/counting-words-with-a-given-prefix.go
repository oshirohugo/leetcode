package countingwordswithagivenprefix

func prefixCount(words []string, pref string) int {
	l := len(pref)
	count := 0
	for _, word := range words {
		if len(word) >= l && word[:l] == pref {
			count++
		}
	}
	return count
}
