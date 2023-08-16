package checkifbinarystringhasatmostonesegmentofones

import "strings"

func checkOnesSegment(s string) bool {
	return !strings.Contains(s, "01")
}

func checkOnesSegmentSlow(s string) bool {
	noMoreOnes := false
	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			noMoreOnes = true
		} else if noMoreOnes {
			return false
		}
	}

	return true
}

// TODO commit
