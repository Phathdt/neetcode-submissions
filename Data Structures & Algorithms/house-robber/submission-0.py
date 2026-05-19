class Solution:
    def rob(self, nums: List[int]) -> int:
        dp = [0] * (len(nums) + 2)

        for i, num in enumerate(nums):
            dp[i + 2] = max(dp[i + 1], dp[i] + num)

        return dp[-1]
