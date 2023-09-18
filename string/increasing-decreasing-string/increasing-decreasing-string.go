package increasingdecreasingstring

func sortString(s string) string {
	count := [26]rune{}
	for _, c := range s {
		count[c-'a']++
	}

	ans := make([]rune, 0, len(s))

	for len(ans) != len(s) {
		for i := 0; i < 26; i++ {
			if count[i] > 0 {
				count[i]--
				ans = append(ans, rune(i+'a'))
			}
		}
		for i := 25; i >= 0; i-- {
			if count[i] > 0 {
				count[i]--
				ans = append(ans, rune(i+'a'))
			}
		}
	}

	return string(ans)
}
