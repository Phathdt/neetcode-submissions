import (
	"cmp"
	"slices"
)


func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int, 0)
	for _, num := range nums {
		count[num] += 1
	}

	ans := make([]int, len(count))
	i := 0
	for num := range count {
		ans[i] = num
		i++
	}

	slices.SortFunc(ans, func(a, b int) int {
		return cmp.Compare(count[b], count[a])
	})

	return ans[:k]
}
