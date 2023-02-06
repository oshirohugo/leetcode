package countthenumberofconsistentstrings

func countConsistentStrings(allowed string, words []string) int {
	lookup := make(map[rune]bool)

	for _, c := range allowed {
		lookup[c] = true
	}
	consistent := len(words)
	for _, w := range words {
		for _, c := range w {
			if !lookup[c] {
				consistent--
				break
			}
		}
	}

	return consistent

}
