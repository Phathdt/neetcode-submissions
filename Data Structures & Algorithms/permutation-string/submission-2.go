func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	counterS1 := toCounter(s1)
	counterS2 := toCounter(s2[:len(s1)])

	if isEqual(counterS1, counterS2) {
		return true
	}

	left := 0

	for right := len(s1); right < len(s2); right++ {
		ch := s2[right]
		counterS2[ch-'a'] += 1
		counterS2[s2[left]-'a'] -= 1

		if isEqual(counterS1, counterS2) {
			return true
		}

		left++

	}
	return false
}

func isEqual(s1 []byte, s2 []byte) bool {
	if len(s1) != len(s2) {
		return false
	}

	for index := range s1 {
		if s1[index] != s2[index] {
			return false
		}
	}

	return true
}
func toCounter(s string) []byte {
	result := make([]byte, 26)

	for i := 0; i < len(s); i++ {
		ch := s[i]

		result[ch-'a'] += 1
	}

	return result
}