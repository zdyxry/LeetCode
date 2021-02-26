# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def nextLargerNodes(self, head: ListNode) -> List[int]:
        l = []
        while head:
            l.append(head.val)
            head = head.next
        res = []
        stack = []
        for i in range(len(l)-1, -1, -1):
            while stack and stack[-1] <= l[i]:
                stack.pop()
            if stack:
                res.append(stack[-1])
            else:
                res.append(0)
            stack.append(l[i])
        return res[::-1]