func isPalindrome(s string) bool {
	cleaned := make([]rune, 0)

	for _, c := range s {
		if isAlphaDigit(c) {
			cleaned = append(cleaned, unicode.ToLower(c))
		}
	}

	left := 0
	right := len(cleaned) - 1

	for left < right {
		if cleaned[left] != cleaned[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphaDigit(c rune) bool {
	return unicode.IsDigit(c) || unicode.IsLetter(c)
}