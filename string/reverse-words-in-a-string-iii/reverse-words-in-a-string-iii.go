package reversewordsinastringiii

func reverseWords(s string) string {
	reversed := []rune(s)

	b := 0
	var e int
	for e = 0; e < len(reversed); e++ {
		if reversed[e] == ' ' {
			reverse(reversed, b, e-1)
			b = e + 1
		}
	}
	reverse(reversed, b, e-1)

	return string(reversed)
}

func reverse(s []rune, b, e int) {
	j := e
	i := b
	for i < j {
		tmp := s[i]
		s[i] = s[j]
		s[j] = tmp
		j--
		i++
	}
}
