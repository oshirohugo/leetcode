package minimizestringlength

func minimizedStringLength(s string) int {
	set := make(map[rune]struct{}, 0)
	for _, c := range s {
		set[c] = struct{}{}
	}

	return len(set)
}
