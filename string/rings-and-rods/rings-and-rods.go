package ringsandrods

var colors = "RGB"

func countPoints(rings string) int {
	lookup := make(map[rune]map[rune]bool)
	r := []rune(rings)

	for i := 0; i < len(r)-1; i += 2 {
		color := r[i]
		rod := r[i+1]

		_, ok := lookup[rod]
		if !ok {
			lookup[rod] = make(map[rune]bool)
		}
		lookup[rod][color] = true
	}
	fullRods := 0
	for rod := '0'; rod <= '9'; rod++ {
		count := 0
		for _, color := range colors {
			if lookup[rod][color] {
				count++
			}
		}
		if count == 3 {
			fullRods++
		}
	}

	return fullRods
}
