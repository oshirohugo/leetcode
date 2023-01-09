package nondecreasingarray

func checkPossibility(nums []int) bool {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			count++
			if count > 1 || nonFixableCase(nums, i) {
				return false
			}

		}
	}
	return true
}

func nonFixableCase(nums []int, i int) bool {
	return isInsideLimits(nums, i) && nums[i-2] > nums[i] && nums[i+1] < nums[i-1]
}

func isInsideLimits(nums []int, i int) bool {
	return i > 1 && i < len(nums)-1
}
