func minSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	ans := len(nums) + 1

	for right := range nums {
		sum += nums[right]

		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	if ans == len(nums)+1 {
		return 0
	}

	return ans
}