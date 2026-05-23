func topKFrequent(nums []int, k int) []int {
	seen := make(map[int]int)

	for _, num := range nums {
		seen[num] += 1
	}

	result := make([]int, len(seen))
	i := 0
	for num := range seen {
		result[i] = num
		i++
	}

	sort.Slice(result, func(i, j int) bool {
		return seen[result[i]] > seen[result[j]]
	})

	return result[:k]
}