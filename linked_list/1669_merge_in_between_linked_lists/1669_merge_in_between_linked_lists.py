# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

        
class Solution:
    def mergeInBetween(self, list1: ListNode, a: int, b: int, list2: ListNode) -> ListNode:
        dummy = ListNode(0)
        dummy.next = list1
        l = dummy
        c = -1
        d1,d2 = None,None
        while l:
            if c == a - 1:
                d1 = l
            if c == b + 1:
                d2 = l
            l = l.next
            c += 1
        if d2 == None:
            d1.next = list2
        else:
            d1.next = list2
            q = list2
            while q.next:
                q = q.next
            q.next = d2
        return dummy.next
