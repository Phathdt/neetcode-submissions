func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int, 0)
	for _, num := range nums {
		count[num] += 1
	}

	bucket := make([][]int, len(nums)+1)
	for num, freq := range count {
		bucket[freq] = append(bucket[freq], num)
	}

	result := make([]int, 0, k)

	for freq := len(nums); freq > 0; freq-- {
		numbers := bucket[freq]

		for _, num := range numbers {
			result = append(result, num)

			if len(result) >= k {
				return result
			}
		}
	}

	return result
}