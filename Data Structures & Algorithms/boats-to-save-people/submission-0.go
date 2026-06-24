
func numRescueBoats(people []int, limit int) int {
	boats := 0
	left, right := 0, len(people)-1

	// Sort the people array in ascending order
	sort.Ints(people)

	for left <= right {
		if people[left]+people[right] <= limit {
			left++
		}

		right--
		boats++
	}

	return boats
}
