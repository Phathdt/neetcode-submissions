func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complete := target - num

		if position, exist := seen[complete]; exist {
			return []int{position, i}
		}

		seen[num] = i
	}
	return []int{}
}
