package make

func minSubarray(nums []int, p int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	target := positiveRemain(sum, p)

	if target == 0 {
		return 0
	}

	lookup := make(map[int]int)
	lookup[0] = -1
	length := len(nums)
	sum = 0
	for j, num := range nums {
		sum += num
		remain := positiveRemain(sum, p)
		key := positiveRemain(remain-target, p)

		if i, ok := lookup[key]; ok {
			currentLength := j - i
			if currentLength < length {
				length = currentLength
			}
		}

		lookup[remain] = j
	}

	if length == len(nums) {
		return -1
	}

	return length

}

func positiveRemain(num, div int) int {
	remain := num % div
	if num < 0 {
		return remain + div
	}
	return remain
}
