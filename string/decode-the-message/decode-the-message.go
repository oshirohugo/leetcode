package decodethemessage

func decodeMessage(key string, message string) string {
	dic := map[rune]rune{
		' ': ' ',
	}

	letter := 'a'
	for _, c := range key {
		_, ok := dic[c]
		if c == ' ' || ok {
			continue
		}

		dic[c] = letter
		letter++
	}

	ans := make([]rune, len(message))

	for i, c := range message {
		ans[i] = dic[c]
	}

	return string(ans)
}
