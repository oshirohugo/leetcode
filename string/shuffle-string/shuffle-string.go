package shufflestring

func restoreString(s string, indices []int) string {
	ans := make([]byte, len(s))
	for i, new := range indices {
		ans[new] = s[i]
	}

	return string(ans)
}
