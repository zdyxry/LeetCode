class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def hasPathSum(self, root, sum):
        if root is None:
            return False
        
        if root.left is None and root.right is None and root.val == sum:
            return True 
        return self.hasPathSum(root.left, sum - root.val) or self.hasPathSum(root.right, sum - root.val)





# test = [5,4,8,11,null,13,4,7,2,null,null,null,1]

root = TreeNode(5)
root.left = TreeNode(4)
root.right = TreeNode(8)
root.left.left = TreeNode(11)
root.left.left.left = TreeNode(7)
root.left.left.right = TreeNode(2)
root.right.left = TreeNode(13)
root.right.right = TreeNode(4)
root.right.right.right = TreeNode(1)

print(Solution().hasPathSum(root, 22))