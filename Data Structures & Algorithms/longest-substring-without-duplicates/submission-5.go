func lengthOfLongestSubstring(s string) int {
    lastSeen := make(map[byte]int)

    longest := 0 
    left := 0 

    for right := 0; right < len(s); right ++ {
        ch := s[right]

        if idx, ok := lastSeen[ch]; ok && idx >= left {
			left = idx + 1
		}

		lastSeen[ch] = right

		longest = max(longest, right-left+1)
    }

    return longest 
}

func max(a, b int) int { 
    if a > b {
        return a 
    }

    return b 
}