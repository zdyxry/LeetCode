class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def isPalindrome(self, head):
        fast = slow = head
        while fast and fast.next:
            fast = fast.next.next
            slow = slow.next
        
        node = None
        while slow:
            nxt = slow.next
            slow.next = node
            node = slow
            slow = nxt
        
        while node:
            if node.val != head.val:
                return False
            node, head = node.next, head.next
        return True


head = ListNode(1)
head.next = ListNode(2)
head.next.next = ListNode(3)
print(Solution().isPalindrome(head))