class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        num_set = set(nums)
        ans = 0

        for num in nums:
            if num - 1 in num_set:
                continue

            curr = num
            leng = 1

            while curr + 1 in num_set:
                curr += 1
                leng += 1

            ans = max(ans, leng)

        return ans
