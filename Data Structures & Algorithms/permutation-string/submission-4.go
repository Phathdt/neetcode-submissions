
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	dictS1 := toDict(s1)
	dictS2 := toDict(s2[:len(s1)])

	if dictS1 == dictS2 {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		dictS2[s2[i]-'a']++
		dictS2[s2[i-len(s1)]-'a']--

		if dictS1 == dictS2 {
			return true
		}
	}

	return false
}

func toDict(s string) [26]int {
	var res [26]int

	for _, char := range s {
		res[char-'a']++
	}

	return res
}