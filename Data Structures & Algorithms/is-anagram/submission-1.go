func isAnagram(s string, t string) bool {
	countS := make(map[rune]int)
	countT := make(map[rune]int)

	for _, c := range s {
		countS[c] += 1
	}
	for _, c := range t {
		countT[c] += 1
	}

	if len(countS) != len(countT) {
		return false
	}

	for k := range countS {
		if countS[k] != countT[k] {
			return false
		}
	}
	return true
}