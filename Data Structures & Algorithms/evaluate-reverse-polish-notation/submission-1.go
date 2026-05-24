
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}

	for _, token := range tokens {
		if operators[token] {
			var a int
			var b int

			stack, b = pop(stack)
			stack, a = pop(stack)

			var newToken int

			switch token {
			case "+":
				newToken = a + b
			case "-":
				newToken = a - b
			case "*":
				newToken = a * b
			case "/":
				newToken = a / b
			}

			stack = append(stack, newToken)
		} else {
			num, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
			stack = append(stack, num)
		}
	}
	return stack[len(stack)-1]
}

func pop(list []int) ([]int, int) {
	tail, list := list[len(list)-1], list[:len(list)-1]
	return list, tail
}