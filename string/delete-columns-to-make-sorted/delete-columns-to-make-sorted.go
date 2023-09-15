package deletecolumnstomakesorted

func minDeletionSize(strs []string) int {
	nOfCols := len(strs[0])
	del := 0
	for col := 0; col < nOfCols; col++ {
		for row := 0; row < len(strs)-1; row++ {
			if strs[row][col] > strs[row+1][col] {
				del++
				break
			}
		}
	}
	return del
}
