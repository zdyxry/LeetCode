class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def reverseList(self, head):
        prev = None
        while head:
            curr = head
            head = head.next
            curr.next = prev
            prev = curr
        return prev


head = ListNode(1)
head.next = ListNode(2)
head.next.next = ListNode(3)
result = Solution().reverseList(head)
print(result.next.next.val)