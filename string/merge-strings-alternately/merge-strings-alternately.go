package mergestringsalternately

func mergeAlternately(word1 string, word2 string) string {
	l1 := len(word1)
	l2 := len(word2)

	i1 := 0
	i2 := 0
	i := 0

	ans := make([]byte, l1+l2)

	for i1 < l1 && i2 < l2 {
		if i%2 == 0 {
			ans[i] = word1[i1]
			i1++
		} else {
			ans[i] = word2[i2]
			i2++
		}
		i++
	}

	if i1 < l1 {
		for i1 < l1 {
			ans[i] = word1[i1]
			i1++
			i++
		}
	}

	if i2 < l2 {
		for i2 < l2 {
			ans[i] = word2[i2]
			i2++
			i++
		}
	}

	return string(ans)
}
