package numberoflinestowritestring

func numberOfLines(widths []int, s string) []int {
	buffer := 100
	lines := 0

	for i := 0; i < len(s); {
		try := buffer - widths[s[i]-'a']

		if try < 0 {
			lines++
			buffer = 100
		} else {
			buffer = try
			i++
		}
	}

	if buffer < 100 {
		lines++
	}

	return []int{
		lines,
		100 - buffer,
	}
}
