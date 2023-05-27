package numberofseniorcitizens

import "strconv"

func countSeniors(details []string) int {
	total := 0
	for _, details := range details {
		if n, _ := strconv.Atoi(details[11:13]); n > 60 {
			total++
		}
	}

	return total
}
