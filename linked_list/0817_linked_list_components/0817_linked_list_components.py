# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def numComponents(self, head: ListNode, G: List[int]) -> int:
        cnt = 0 
        tmp = []
        while head:
            if head.val in G:
                tmp.append(head.val)
            if head.val not in G or not head.next:
                if len(tmp) > 0:
                    cnt += 1
                    tmp = []
            head = head.next
        return cnt
            
            
