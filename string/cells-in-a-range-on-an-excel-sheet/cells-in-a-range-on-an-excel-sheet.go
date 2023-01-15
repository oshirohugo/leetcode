package cellsinarangeonanexcelsheet

func cellsInRange(s string) []string {
	colLow := s[0]
	colHigh := s[3]
	rowLow := s[1]
	rowHigh := s[4]

	ansSize := (colHigh - colLow + 1) * (rowHigh - rowLow + 1)

	ans := make([]string, int(ansSize), int(ansSize))
	cel := make([]byte, 2, 2)

	i := 0
	for col := colLow; col <= colHigh; col++ {
		cel[0] = col
		for row := rowLow; row <= rowHigh; row++ {
			cel[1] = row
			ans[i] = string(cel)
			i++
		}
	}

	return ans
}
