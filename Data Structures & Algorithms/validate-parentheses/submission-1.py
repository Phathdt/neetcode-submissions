class Solution:
    def isValid(self, s: str) -> bool:
        pair = {"}": "{", ")": "(", "]": "["}

        queue = []

        for i in range(len(s)):
            char = s[i]

            if char in pair:
                if len(queue) == 0:
                    return False 
                last = queue.pop()

                if last != pair[char]:
                    return False

            else:
                queue.append(char)

        return len(queue) == 0
