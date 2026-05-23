func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1 
	for i, num := range nums {
		result[i] = prefix 
		prefix = prefix * num 
	}

	suffix := 1
	for i := len(nums)-1; i >= 0; i-- { 
		result[i] *= suffix
		suffix = suffix * nums[i]
	}

	return result 
}

// a 	b 	c 
// 1 	a 	ab 
// bc  c 	1