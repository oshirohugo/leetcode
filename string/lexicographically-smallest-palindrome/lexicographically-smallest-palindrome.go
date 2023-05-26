package lexicographicallysmallestpalindrome

func makeSmallestPalindrome(s string) string {
	i := 0
	j := len(s) - 1
	w := []byte(s)
	for i < j {
		if w[i] < w[j] {
			w[j] = w[i]
		}
		if w[i] > w[j] {
			w[i] = w[j]
		}
		i++
		j--
	}

	return string(w)
}
