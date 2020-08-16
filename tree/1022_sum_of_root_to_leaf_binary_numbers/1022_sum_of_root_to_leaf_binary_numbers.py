# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

        
class Solution:
    def sumRootToLeaf(self, root: TreeNode) -> int:
        return self.helper(root, 0)
    def helper(self, root, base):
        if root == None:
            return 0
        if root.left == None and root.right == None:
            return 2 * base + root.val
        base = 2 * base + root.val
        return self.helper(root.left, base) + self.helper(root.right, base)