package destinationcity

func destCity(paths [][]string) string {
	cities := map[string]int{}
	for _, path := range paths {
		cities[path[0]]++
		cities[path[1]]--
	}
	for city, value := range cities {
		if value == -1 {
			return city
		}
	}

	return ""
}
