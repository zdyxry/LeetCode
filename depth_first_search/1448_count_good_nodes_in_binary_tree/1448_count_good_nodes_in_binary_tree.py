# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def goodNodes(self, root: TreeNode) -> int:
        self.res = 0

        def dfs(node, curmx):
            if not node:
                return
            if node.val >= curmx:
                self.res += 1
            curmx = max(curmx, node.val)
            dfs(node.left, curmx)
            dfs(node.right, curmx)

        dfs(root, -float('inf'))
        return self.res
