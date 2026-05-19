# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        if not head or not head.next:
            return

        slow, fast = head, head
        prev = None

        while fast and fast.next:
            prev = slow
            slow = slow.next
            fast = fast.next.next

        prev.next = None

        first = head
        second = self.reverse(slow)

        while first and second:
            tmp1 = first.next
            tmp2 = second.next

            first.next = second
            if not tmp1:
                break

            second.next = tmp1

            first = tmp1
            second = tmp2

    def reverse(self, head):
        prev = None
        cur = head

        while cur:
            next = cur.next
            cur.next = prev
            prev = cur
            cur = next

        return prev