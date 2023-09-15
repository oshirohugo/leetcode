package deletecolumnstomakesorted

func minDeletionSize(strs []string) int {
	nOfCols := len(strs[0])
	del := 0
	for i := 0; i < nOfCols; i++ {
		prev := strs[0][i]
		for j := 1; j < len(strs); j++ {
			curr := strs[j][i]
			if prev > curr {
				del++
				break
			}
			prev = curr
		}
	}
	return del
}
