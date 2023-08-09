package checkifallcharactershaveequalnumberofoccurrences

func areOccurrencesEqual(s string) bool {
	count := [27]int{}

	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}

	f := count[s[0]-'a']

	for _, v := range count {
		if v != 0 && v != f {
			return false
		}
	}

	return true
}

func areOccurrencesEqualMoreSpace(s string) bool {
	count := make(map[rune]int)
	var f int

	for _, c := range s {
		if _, ok := count[c]; ok {
			count[c]++
		} else {
			count[c] = 1
		}
		f = count[c]
	}

	for _, v := range count {
		if v != f {
			return false
		}
	}

	return true
}
