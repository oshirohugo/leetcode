package kthdistinctstringinanarray

func kthDistinct(arr []string, k int) string {
	count := map[string]int{}

	for _, a := range arr {
		count[a]++
	}

	unique := 1
	for _, a := range arr {
		if count[a] == 1 {
			if unique == k {
				return a
			}
			unique++
		}
	}

	return ""
}
