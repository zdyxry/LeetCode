
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def rotateRight(self, head, k):
        if not head or not head.next:
            return head

        n, cur = 1, head
        while cur.next:
            cur = cur.next
            n += 1
        cur.next = head

        cur, tail = head, cur
        for _ in xrange(n - k % n):
            tail = cur
            cur = cur.next
        tail.next = None 

        return cur 


head = ListNode(1)
head.next = ListNode(2)
head.next.next = ListNode(3)
res = Solution().rotateRight(head, 2)
print(res.val)