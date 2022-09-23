func subarraySum(nums []int, k int) int {
	prefix := make([]int, len(nums))
	prefix[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefix[i] = nums[i] + prefix[i-1]
	}

	ans := 0
	sum := make(map[int]int)
	for _, p := range prefix {
		if p == k {
			ans++
		}

		diff := p - k

		if pp, ok := sum[diff]; ok {
			ans += pp
		}

		sum[p]++
	}

	return ans
}
