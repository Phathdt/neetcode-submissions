from functools import cache


class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        n = len(nums)

        @cache
        def dfs(i):
            ans = 1

            for j in range(i + 1, n):
                if nums[j] > nums[i]:
                    ans = max(ans, 1 + dfs(j))

            return ans

        res = 0

        for i in range(n):
            res = max(res, dfs(i))

        return res
