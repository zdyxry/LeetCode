# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

        
class Solution:
    def pseudoPalindromicPaths (self, root: TreeNode) -> int:
        s = set()

        def dfs(root):

            if not root: return 0

            if root.val in s:
                s.discard(root.val)
            else:
                s.add(root.val)

            res = dfs(root.left) + dfs(root.right)

            if not root.left and not root.right:
                res += len(s)<=1

            if root.val in s:
                s.discard(root.val)
            else:
                s.add(root.val)

            return res

        return dfs(root)