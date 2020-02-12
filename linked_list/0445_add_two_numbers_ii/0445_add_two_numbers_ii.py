class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def addTwoNumbers(self, l1, l2):
        s1, s2 = [], []
        while l1:
            s1.append(l1.val)
            l1 = l1.next
        while l2:
            s2.append(l2.val)
            l2 = l2.next

        tail = ListNode(-1)
        s = 0

        while s1 or s2:
            a, b = s1.pop() if s1 else 0, s2.pop() if s2 else 0
            s += a + b
            tail.val = s % 10
            s /= 10
            head = ListNode(s)
            head.next = tail
            tail = head

        return tail if tail.val else tail.next