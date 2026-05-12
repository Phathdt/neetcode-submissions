class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        l, r = 1, max(piles)

        def can_finish(k):
            hours = 0

            for pile in piles:
                hours += (pile + k - 1) // k

            return hours <= h

        ans = r

        while l <= r:
            mid = l + (r - l) // 2

            if can_finish(mid):
                ans = mid
                r = mid - 1
            else:
                l = mid + 1

        return ans