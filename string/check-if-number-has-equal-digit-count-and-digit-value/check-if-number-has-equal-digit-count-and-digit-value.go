package checkifnumberhasequaldigitcountanddigitvalue

func digitCount(num string) bool {
	var count [10]int

	for _, c := range num {
		count[int(c-'0')]++
	}

	for i := 0; i < len(num); i++ {
		if count[i] != int(num[i]-'0') {
			return false
		}
	}
	return true
}
