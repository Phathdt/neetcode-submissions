class Solution:
    def dailyTemperatures(self, temps: List[int]) -> List[int]:
        n = len(temps)
        answer = [0] * n
        stack = [] 

        for i, t in enumerate(temps):
            while stack and t > temps[stack[-1]]:
                idx = stack.pop()
                answer[idx] = i - idx

            stack.append(i)
        return answer