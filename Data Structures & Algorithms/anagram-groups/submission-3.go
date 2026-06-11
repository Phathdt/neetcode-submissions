import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	group := make(map[string][]string)

	for _, str := range strs {
		key := generateKey(str)
		group[key] = append(group[key], str)
	}

	ans := make([][]string, 0)
	for _, groupStr := range group {
		ans = append(ans, groupStr)
	}

	return ans
}

func generateKey(str string) string {
	chars := []byte(str)
	slices.Sort(chars)

	return string(chars)
}
