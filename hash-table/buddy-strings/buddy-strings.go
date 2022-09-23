package buddystrings

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	lookup := [int('z' - 'a')]int{}
	repeat := false
	diffs := 0
	indexes := [2]int{}

	for i, _ := range s {
		if s[i] != goal[i] {
			diffs++
			if diffs > 2 {
				return false
			}
			indexes[diffs-1] = i
		}
		if lookup[int(s[i]-'a')] > 0 {
			repeat = true
		}
		lookup[int(s[i]-'a')]++
	}

	if diffs == 0 {
		return repeat
	}

	if diffs == 2 {
		return s[indexes[0]] == goal[indexes[1]] && s[indexes[1]] == goal[indexes[0]]
	}

	return false
}
