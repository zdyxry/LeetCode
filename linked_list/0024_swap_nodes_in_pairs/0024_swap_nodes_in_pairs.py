class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def swapPairs(self, head):
        thead = ListNode(-1)
        thead.next = head
        c = thead
        while c.next and c.next.next:
            a,b = c.next, c.next.next
            c.next, a.next = b, b.next
            b.next = a
            c = c.next.next
        return thead.next


a = ListNode(1)
a.next = ListNode(2)
a.next.next = ListNode(3)

res = Solution().swapPairs(a)
print(res.val)