func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int)

	for i, v := range nums {
		if j, exist := seen[v]; exist {
			if i-j <= k {
				return true
			}
		}

		seen[v] = i
	}

	return false
}
