package maximumrepeatingsubstring

func maxRepeating(sequence string, word string) int {
	wordLen := len(word)
	sequenceLen := len(sequence)
	k := 0
	i := 0
	for sequenceLen-i >= wordLen {
		if sequence[i:i+wordLen] == word {
			k++
			i += wordLen
		} else {
			i++
		}
	}

	return k
}
