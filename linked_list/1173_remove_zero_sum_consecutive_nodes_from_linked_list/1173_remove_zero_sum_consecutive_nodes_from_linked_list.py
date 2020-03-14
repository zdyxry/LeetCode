class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def removeZeroSumSublists(self, head):
        dummy = ListNode(0)
        dummy.next = head
        prefix = 0
        d = {0:dummy}
        while head:
            prefix += head.val
            d[prefix] = head
            head = head.next
        
        head = dummy
        prefix = 0
        while head:
            prefix += head.val
            head.next = d[prefix].next
            head = head.next

        return dummy.next