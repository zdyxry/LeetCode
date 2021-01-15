# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    def swapNodes(self, head: ListNode, k: int) -> ListNode:
        slow = fast = tmp = head
        for _ in range(k-1):
            fast, tmp = fast.next, tmp.next
        while fast.next:
            fast = fast.next
            slow = slow.next
        tmp.val, slow.val = slow.val, tmp.val
        return head
