
func isValid(s string) bool {
	stack := make([]byte, 0)
	dict := map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}

	for index := range s {
		ch := s[index]

		if _, exist := dict[ch]; exist {
			if len(stack) == 0 {
				return false
			}

			var last byte
			stack, last = pop(stack)

			if dict[ch] != last {
				return false
			}
		} else {
			stack = append(stack, ch)
		}

	}
	return len(stack) == 0
}

func pop(arr []byte) ([]byte, byte) {
	last := arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	return arr, last
}