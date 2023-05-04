package maximumrepeatingsubstring

import "strings"

func maxRepeating(sequence string, word string) int {
	k := 0
	acc := word
	for strings.Contains(sequence, acc) {
		k++
		acc += word
	}
	return k
}
