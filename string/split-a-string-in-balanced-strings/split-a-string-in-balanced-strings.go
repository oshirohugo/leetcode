package splitastringinbalancedstrings

func balancedStringSplit(s string) int {
	res := 0

	count := 0
	for _, c := range s {
		if c == 'L' {
			count += 1
		} else {
			count -= 1
		}

		if count == 0 {
			res++
		}
	}

	return res
}
