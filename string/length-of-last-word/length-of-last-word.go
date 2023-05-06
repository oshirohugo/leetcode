package lengthoflastword

import "strings"

func lengthOfLastWord(s string) int {
	tail := len(s) - 1
	for tail >= 0 && s[tail] == ' ' {
		tail--
	}
	ans := 0
	for tail >= 0 && s[tail] != ' ' {
		ans++
		tail--
	}

	return ans
}

func easyLengthOfLastWord(s string) int {
	w := strings.Fields(s)
	return len(w[len(w)-1])
}
