import "slices"
func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	ans := make([][]int, 0)

	n := len(nums)

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := n - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return ans
}