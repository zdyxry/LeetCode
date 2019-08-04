# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def invertTree(self, root):
        if root is not None:
            root.left , root.right = self.invertTree(root.right), self.invertTree(root.left)
        
        return root

root = TreeNode(1)
root.left = TreeNode(2)
root.right = TreeNode(3)
res = Solution().invertTree(root)
print(res.right.val)