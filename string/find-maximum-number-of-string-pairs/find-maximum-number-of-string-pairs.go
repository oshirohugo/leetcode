package findmaximumnumberofstringpairs

func maximumNumberOfStringPairs(words []string) int {
	lookup := map[string]struct{}{}
	pair := 0
	for _, word := range words {
		if _, ok := lookup[string([]byte{word[1], word[0]})]; ok {
			pair++
		}
		lookup[word] = struct{}{}
	}
	return pair
}
