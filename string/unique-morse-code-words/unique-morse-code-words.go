package uniquemorsecodewords

import "strings"

func uniqueMorseRepresentations(words []string) int {
	unique := make(map[string]bool)

	for _, word := range words {
		c := getCode(word)
		if _, ok := unique[c]; !ok {
			unique[c] = true
		}
	}

	return len(unique)
}

var code = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func getCode(word string) string {
	a := rune('a')
	var ans []string
	for _, c := range word {
		ans = append(ans, code[c-a])
	}

	return strings.Join(ans, "")
}
