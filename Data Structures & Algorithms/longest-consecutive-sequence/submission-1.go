func longestConsecutive(nums []int) int {
	num_set := make(map[int]bool)

	for _, num := range nums {
		num_set[num] = true
	}

	ans := 0

	for num := range num_set {
		if _, exist := num_set[num-1]; exist {
			continue
		}

		cur := num
		leng := 1

		for num_set[cur+1] {
			cur += 1
			leng += 1
		}

		ans = max(ans, leng)
	}

	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}