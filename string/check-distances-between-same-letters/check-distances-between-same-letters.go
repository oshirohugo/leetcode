package checkdistancesbetweensameletters

func checkDistances(s string, distance []int) bool {
	lookup := make(map[rune]int)
	for i, c := range s {
		if old, ok := lookup[c]; ok {
			if (i - old - 1) != distance[c-'a'] {
				return false
			}
		} else {
			lookup[c] = i
		}
	}

	return true
}
