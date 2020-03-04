# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def isSubPath(self, head, root):
        if not root:
            return False

        return self.dfs(head, root) or self.isSubPath(head, root.left) or self.isSubPath(head, root.right)

    def dfs(self, head, root):
        if not head:
            return True
        if not root:
            return False
        if head.val != root.val:
            return False
        else:
            return self.dfs(head.next, root.left) or self.dfs(head.next, root.right)


            