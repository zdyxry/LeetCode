class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def middleNode(self, head):
        slow, fast = head, head
        while fast and fast.next:
            fast = fast.next.next
            slow = slow.next
        return slow

head = ListNode(1)
head.next = ListNode(2)
head.next.next = ListNode(3)
head.next.next.next = ListNode(4)

res = Solution().middleNode(head)
print(res.val)