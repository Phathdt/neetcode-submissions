from collections import Counter


class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        if len(s1) > len(s2):
            return False

        need = Counter(s1)
        window = Counter()
        l = 0

        for r in range(len(s2)):
            window[s2[r]] += 1

            if r - l + 1 > len(s1):
                window[s2[l]] -= 1

                l += 1

            if need == window:
                return True

        return False
