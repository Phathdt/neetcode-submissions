func dailyTemperatures(temps []int) []int {
	n := len(temps)
	ans := make([]int, n)
	stack := make([]int, 0)

	for i, temp := range temps {
		for len(stack) > 0 && temp > temps[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ans[top] = i - top
		}
		stack = append(stack, i)
	}
	return ans
}