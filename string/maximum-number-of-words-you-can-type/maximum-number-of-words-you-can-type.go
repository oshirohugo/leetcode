package maximumnumberofwordsyoucantype

func canBeTypedWords(text string, brokenLetters string) int {
	lookup := make(map[byte]bool)

	for i := 0; i < len(brokenLetters); i++ {
		lookup[brokenLetters[i]] = true
	}

	ignore := false
	count := 0
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			if ignore == false {
				count++
			}
			ignore = false
		}
		if lookup[text[i]] {
			ignore = true
		}
	}

	if ignore == false {
		count++
	}

	return count
}
