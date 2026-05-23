import (
	"slices"
)


func groupAnagrams(strs []string) [][]string {
	group := make(map[string][]string)

	for _, str := range strs {
		key := sortString(str)

		group[key] = append(group[key], str)
	}

	result := make([][]string, 0)
	for _, strs := range group {
		result = append(result, strs)
	}
	return result
}

func sortString(s string) string {
	chars := []byte(s)

	slices.Sort(chars)

	return string(chars)
}