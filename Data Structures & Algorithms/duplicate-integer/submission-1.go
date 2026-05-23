func hasDuplicate(nums []int) bool {
    seen := make(map[int]bool)

    for _, value := range nums {
        if _, exist := seen[value]; exist {
            return true
        }

        seen[value] = true
    }

    return false
}
