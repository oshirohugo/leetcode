package countasterisks

func countAsterisks(s string) int {
	counting := true
	asterisks := 0
	for _, c := range s {
		if c == '|' {
			if counting {
				counting = false
			} else {
				counting = true
			}
		}

		if c == '*' && counting {
			asterisks++
		}
	}

	return asterisks
}
