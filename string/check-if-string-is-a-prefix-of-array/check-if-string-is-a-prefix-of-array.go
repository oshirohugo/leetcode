package checkifstringisaprefixofarray

func isPrefixString(s string, words []string) bool {
	check := ""
	for _, c := range words {
		check += c
		if check == s {
			return true
		}
		if len(check) > len(s) {
			return false
		}
	}

	return false
}
