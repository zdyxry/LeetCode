class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def minDepth(self,root):
        if root is None:
            return 0 
        
        if root.left and root.right:
            return min(self.minDepth(root.left), self.minDepth(root.right)) + 1 
        else:
            return max(self.minDepth(root.left), self.minDepth(root.right)) + 1 



root = TreeNode(1)
root.left = TreeNode(2)
root.right = TreeNode(3)
root.right.right= TreeNode(4)

print(Solution().minDepth(root))