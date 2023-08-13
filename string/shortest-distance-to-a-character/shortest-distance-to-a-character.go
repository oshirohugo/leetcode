package shortestdistancetoacharacter

func shortestToChar(s string, c byte) []int {
	dist := make([]int, len(s))
	n := len(s)
	pos := -n
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			pos = i
		}
		dist[i] = i - pos
	}
	for i := pos - 1; i >= 0; i-- {
		if s[i] == c {
			pos = i
		}
		dist[i] = min(dist[i], pos-i)
	}

	return dist
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func shortestToCharSlow(s string, c byte) []int {
	var index []int
	dist := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == c {
			index = append(index, i)
		}
	}

	currentIndex := index[0]
	nextIndex := currentIndex
	j := 0
	if len(index) > 1 {
		nextIndex = index[1]
		j = 1
	}

	for i := 0; i < len(s); i++ {
		currentDist := distance(i, currentIndex)
		nextDist := distance(i, nextIndex)
		if currentDist < nextDist {
			dist[i] = currentDist
		} else {
			dist[i] = nextDist
			currentIndex = nextIndex
			if j < len(index)-1 {
				j++
				nextIndex = index[j]
			}
		}
	}

	return dist
}

func distance(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
