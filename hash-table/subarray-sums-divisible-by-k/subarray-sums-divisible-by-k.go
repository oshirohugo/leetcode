func subarraysDivByK(nums []int, k int) int {
	freq := make([]int, k)
	sum := 0
	res := 0
	freq[0] = 1

	for _, num := range nums {
		sum += num
		remainder := sum % k

		if remainder < 0 {
			remainder += k
		}

		res += freq[remainder]
		freq[remainder]++
	}

	return res
}
