package redistributecharacterstomakeallstringsequal

func makeEqual(words []string) bool {
	var count [26]int

	for _, word := range words {
		for _, c := range word {
			count[c-'a']++
		}
	}

	for _, c := range count {
		if c%len(words) != 0 {
			return false
		}
	}

	return true
}
