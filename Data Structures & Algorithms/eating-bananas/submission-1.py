class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        l, r = 1, max(piles)

        answer = l

        def can_finish(k):
            hours = 0

            for pile in piles:
                hours += (pile + k - 1) // k

            return hours <= h

        while l <= r:
            m = l + (r - l) // 2

            if can_finish(m):
                answer = m
                r = m - 1
            else:
                l = m + 1

        return answer
