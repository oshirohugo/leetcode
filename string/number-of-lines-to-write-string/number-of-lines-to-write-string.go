package numberoflinestowritestring

func numberOfLines(widths []int, s string) []int {
	buffer := 100
	lines := 0

	for i := 0; i < len(s); {
		try := buffer - widths[s[i]-'a']
		if try <= 0 {
			lines++
			buffer = 100
			if try == 0 {
				i++
			}
		} else {
			buffer = try
			i++
		}
	}

	last := buffer
	if buffer < 100 {
		lines++
		last = 100 - last
	}

	return []int{
		lines,
		last,
	}
}
