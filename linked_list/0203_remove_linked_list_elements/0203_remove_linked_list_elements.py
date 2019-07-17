class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def removeElements(self, head, val):
        dummy = ListNode(0)
        dummy.next = head
        pre = dummy
        while head:
            if head.val == val:
                pre.next, head = head.next, head.next
            else:
                pre, head = head, head.next
        return dummy.next


head = ListNode(1)
head.next = ListNode(2)
head.next.next = ListNode(3)
head.next.next.next = ListNode(3)

result = Solution().removeElements(head, 2)

while result:
    print result.val
    result = result.next