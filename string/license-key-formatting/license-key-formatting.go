package licensekeyformatting

import "strings"

func licenseKeyFormatting(s string, k int) string {
	l := len(s)
	var out []byte

	s = strings.ToUpper(s)

	counter := 0
	for i := l - 1; i >= 0; {
		if counter == k {
			out = prepend(out, '-')
			counter = 0
			continue
		}
		if s[i] != '-' {
			out = prepend(out, s[i])
			counter++
		}
		i--
	}

	return strings.Trim(string(out), "-")
}

func prepend(x []byte, y byte) []byte {
	x = append(x, 0)
	copy(x[1:], x)
	x[0] = y
	return x
}
