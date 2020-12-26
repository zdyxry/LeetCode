# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    res = 0
    def maxAncestorDiff(self,root:TreeNode)->int:
        if not root: return 0
        self.dfs(root,root.val, root.val)
        return self.res
    def dfs(self,root, maxValue, minValue):
        maxValue = max(root.val, maxValue)
        minValue = min(root.val, minValue)
        if root.left is None and root.right is None:
            self.res = max(self.res, abs(maxValue - minValue))
        if root.left:
            self.dfs(root.left,maxValue,minValue)
        if root.right:
            self.dfs(root.right,maxValue,minValue)
