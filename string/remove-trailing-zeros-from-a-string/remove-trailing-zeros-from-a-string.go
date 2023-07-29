package removetrailingzerosfromastring

import "strings"

func removeTrailingZeros(num string) string {
	ei := len(num) - 1

	for {
		if num[ei] != '0' {
			break
		}
		ei--
	}

	return num[:ei+1]
}

func removeTrailingZerosEasy(num string) string {
	return strings.TrimRight(num, "0")
}
