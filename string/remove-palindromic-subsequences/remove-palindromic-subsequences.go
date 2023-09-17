package removepalindromicsubsequences

func removePalindromeSub(s string) int {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return 2
		}
	}

	return 1
}
