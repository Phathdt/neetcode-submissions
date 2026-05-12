class Solution:
    def dailyTemperatures(self, temps: List[int]) -> List[int]:
        ans = [0] * len(temps)
        stack = []

        for i, temp in enumerate(temps):
            while stack and temp > temps[stack[-1]]:
                idx = stack.pop()
                ans[idx] = i - idx

            stack.append(i)
        return ans
