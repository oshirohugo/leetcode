package longercontiguoussegmentsofonesthanzeros

func checkZeroOnes(s string) bool {
	zeros := 0
	maxZeros := -1
	ones := 0
	maxOnes := -1

	for _, c := range s {
		if c == '0' {
			ones = 0
			zeros++
		} else {
			zeros = 0
			ones++
		}
		maxZeros = max(maxZeros, zeros)
		maxOnes = max(maxOnes, ones)
	}

	return maxOnes > maxZeros
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
