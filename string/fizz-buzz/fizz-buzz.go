package fizzbuzz

import "strconv"

func fizzBuzz(n int) []string {
	answer := make([]string, n)
	for i := 0; i < n; i++ {
		c := i + 1
		ans := ""
		if c%3 == 0 {
			ans += "Fizz"
		}
		if c%5 == 0 {
			ans += "Buzz"
		}
		if ans == "" {
			ans = strconv.Itoa(c)
		}
		answer[i] = ans
	}

	return answer
}
