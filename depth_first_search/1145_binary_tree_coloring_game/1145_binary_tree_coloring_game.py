# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

        
class Solution:
    def btreeGameWinningMove(self, root: TreeNode, n: int, x: int) -> bool:
        self.red_left, self.red_right = 0, 0
        
        def dfs(root):
            if not root:
                return 0
            left = dfs(root.left)
            right = dfs(root.right)
            if root.val == x:
                self.red_left = left
                self.red_right = right
            return left + right + 1
        
        dfs(root)
        parent = n - self.red_left - self.red_right - 1
        judge = [parent, self.red_left, self.red_right]
        return any([j > n // 2 for j in judge])
