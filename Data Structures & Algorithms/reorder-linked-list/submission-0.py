# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        if not head or not head.next:
            return

        # 1. find middle
        slow, fast = head, head
        prev = None

        while fast and fast.next:
            prev = slow
            slow = slow.next
            fast = fast.next.next

        # cut list
        prev.next = None

        # 2. reverse second half
        second = self.reverse(slow)

        # 3. merge
        first = head

        while first and second:
            tmp1 = first.next
            tmp2 = second.next

            first.next = second

            # nếu first hết thì stop
            if not tmp1:
                break

            second.next = tmp1

            first = tmp1
            second = tmp2

    def reverse(self, head):
        prev = None
        cur = head

        while cur:
            nxt = cur.next
            cur.next = prev
            prev = cur
            cur = nxt

        return prev
