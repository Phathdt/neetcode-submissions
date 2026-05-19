class Solution:
    def combinationSum(self, nums: List[int], target: int) -> List[List[int]]:
        ans = []

        def backtrack(i, remain, path):
            if remain == 0:
                ans.append(path)
                return

            if remain < 0 or i >= len(nums):
                return

            backtrack(i, remain - nums[i], path + [nums[i]])
            backtrack(i + 1, remain, path)

        backtrack(0, target, [])

        return ans
