package sortthepeople

import "sort"

func sortPeople(names []string, heights []int) []string {
	lookup := make(map[int]string)

	for i := 0; i < len(names); i++ {
		lookup[heights[i]] = names[i]
	}

	sort.Ints(heights)

	j := len(names) - 1
	for i := 0; i < len(names); i++ {
		names[i] = lookup[heights[j-i]]
	}

	return names
}
