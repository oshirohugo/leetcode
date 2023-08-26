package detectcapital

func detectCapitalUse(word string) bool {
	nOfCaps := 0
	for i := 0; i < len(word); i++ {
		if word[i] < 'a' {
			nOfCaps++
		}
	}

	return (nOfCaps == len(word)) || (nOfCaps == 1 && word[0] < 'a') || (nOfCaps == 0)
}

func detectCapitalUseComplex(word string) bool {
	if word[0] < 'a' {
		nOfCaps := 0
		for i := 1; i < len(word); i++ {
			if word[i] < 'a' {
				nOfCaps++
			}
		}
		return nOfCaps == (len(word)-1) || nOfCaps == 0
	}

	for i := 1; i < len(word); i++ {
		if word[i] < 'a' {
			return false
		}
	}

	return true
}
