package maximumvalueofastringinanarray

import "strconv"

func maximumValue(strs []string) int {
	max := -1
	for _, s := range strs {
		value, err := strconv.Atoi(s)
		if err != nil {
			value = len(s)
		}

		if value > max {
			max = value
		}
	}
	return max
}

func maximumValuePedantic(strs []string) int {
	max := -1
	for _, s := range strs {
		value := 0
		base := 1
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] < '0' || s[i] > '9' {
				value = len(s)
				break
			}
			value += (base * int(s[i]-'0'))
			base *= 10
		}
		if value > max {
			max = value
		}
	}
	return max
}
