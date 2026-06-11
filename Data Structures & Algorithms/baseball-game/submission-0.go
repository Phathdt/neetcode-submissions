func calPoints(operations []string) int {
	queue := make([]int, 0)

	for _, ch := range operations {
		switch ch {
		case "C":
			queue = queue[:len(queue)-1]
		case "D":
			last := queue[len(queue)-1]
			queue = append(queue, last*2)
		case "+":
			next := queue[len(queue)-2]
			last := queue[len(queue)-1]
			queue = append(queue, next+last)
		default:
			v, err := strconv.Atoi(ch)
			if err != nil {
				panic(err)
			}

			queue = append(queue, v)
		}
	}

	return sum(queue)
}

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}
