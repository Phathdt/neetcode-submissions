func lengthOfLongestSubstring(s string) int {
    seen := make(map[byte]bool)

    strs := []byte(s)

    left := 0
    longest := 0 

    for right := 0; right < len(s); right++{
        for seen[strs[right]] {
            seen[strs[left]] = false
            left ++
        }
        seen[strs[right]] = true 
        longest = max(longest, right - left + 1)
    } 

    return longest 
}

func max(a, b int) int { 
    if a > b {
        return a 
    }

    return b 
}