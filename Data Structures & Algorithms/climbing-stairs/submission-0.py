class Solution:
    def climbStairs(self, n: int) -> int:
        dp = [0] * (n + 2)
        dp[1] = 1

        for i in range(n):
            dp[i + 2] = dp[i + 1] + dp[i]

        print(dp)
        return dp[-1]
