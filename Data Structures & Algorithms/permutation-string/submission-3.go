
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	counterS1 := toCounter(s1)
	counterS2 := toCounter(s2[:len(s1)])

	if counterS1 == counterS2 {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		counterS2[s2[i]-'a'] += 1
		counterS2[s2[i-len(s1)]-'a'] -= 1

		if counterS1 == counterS2 {
			return true
		}
	}
	return false
}

func toCounter(s string) [26]int {
	var data [26]int

	for i := range s {
		data[s[i]-'a'] += 1
	}
	return data
}