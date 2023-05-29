package mostcommonword

import (
	"regexp"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	re := regexp.MustCompile(`[!?',;.]`)
	paragraph = re.ReplaceAllString(paragraph, " ")
	paragraph = strings.ToLower(paragraph)

	ban := make(map[string]bool)
	for _, b := range banned {
		ban[b] = true
	}

	ans := ""
	max := -1
	lookup := make(map[string]int)
	for _, word := range strings.Fields(paragraph) {
		if ban[word] {
			continue
		}
		lookup[word]++
		if lookup[word] > max {
			max = lookup[word]
			ans = word
		}

	}

	return ans

}
