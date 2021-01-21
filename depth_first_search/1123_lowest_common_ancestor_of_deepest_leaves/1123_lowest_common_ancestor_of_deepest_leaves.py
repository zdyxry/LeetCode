# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

        
class Solution:
    def lcaDeepestLeaves(self, root: TreeNode) -> TreeNode:
        def dfs(root):
            if not root:
                return None, 0
            lr, ld = dfs(root.left)
            rr, rd = dfs(root.right)
            if ld > rd:
                return lr, ld + 1
            elif ld < rd:
                return rr, rd + 1
            else:
                return root, ld + 1
        ans, h = dfs(root)
        return ans
