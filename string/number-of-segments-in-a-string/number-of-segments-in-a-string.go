package numberofsegmentsinastring

func countSegments(s string) int {
	counter := 0
	for i, c := range s {
		if c != ' ' && (i == 0 || s[i-1] == ' ') {
			counter++
		}
	}

	return counter
}
