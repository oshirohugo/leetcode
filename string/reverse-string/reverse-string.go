package reversestring

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1
	for i < j {
		aux := s[i]
		s[i] = s[j]
		s[j] = aux
		i++
		j--
	}
}
