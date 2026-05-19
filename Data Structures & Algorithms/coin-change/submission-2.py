from functools import cache


class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        @cache
        def dfs(remain):
            if remain == 0:
                return 0

            if remain < 0:
                return float("inf")

            ans = float("inf")
            for coin in coins:
                res = dfs(remain - coin)
                ans = min(ans, res + 1)

            return ans

        result = dfs(amount)

        if result == float("inf"):
            return -1

        return result
