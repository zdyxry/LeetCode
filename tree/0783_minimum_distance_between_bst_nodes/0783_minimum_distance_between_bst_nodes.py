# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def minDiffInBST(self, root):
        def dfs(node):
            if not node:
                return 
            dfs(node.left)
            self.result = min(self.result, node.val - self.prev)
            self.prev = node.val
            dfs(node.right)
            
        self.prev = float('-inf')
        self.result = float('inf')
        dfs(root)
        return self.result

root = TreeNode(4)
root.right = TreeNode(6)
root.left = TreeNode(2)
root.left.left = TreeNode(1)
root.left.right = TreeNode(3)

res = Solution().minDiffInBST(root)
print(res)
