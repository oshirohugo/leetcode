package maximumnumberofwordsfoundinsentences

import "strings"

func mostWordsFound(sentences []string) int {
	max := -1
	for _, sentence := range sentences {
		count := len(strings.Fields(sentence))
		if count > max {
			max = count
		}
	}
	return max
}
