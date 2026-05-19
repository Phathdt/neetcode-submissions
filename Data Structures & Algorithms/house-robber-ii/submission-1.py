class Solution:
    def rob1(self, nums: List[int]) -> int:
        dp = [0] * (len(nums) + 2)

        for i, num in enumerate(nums):
            dp[i + 2] = max(dp[i + 1], dp[i] + num)

        return dp[-1]

    def rob(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return nums[0]
            
        return max(self.rob1(nums[1:]), self.rob1(nums[:-1]))
