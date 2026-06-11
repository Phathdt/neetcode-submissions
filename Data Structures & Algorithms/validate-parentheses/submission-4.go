func isValid(s string) bool {
	stack := make([]rune, 0)
	open := map[rune]bool{
		'{': true,
		'[': true,
		'(': true,
	}

	for _, ch := range s {
		if _, exist := open[ch]; exist {
			stack = append(stack, ch)
		} else {
            if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]

			switch last {
			case '{':
				if ch != '}' {
					return false
				}
			case '(':
				if ch != ')' {
					return false
				}
			case '[':
				if ch != ']' {
					return false
				}
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
