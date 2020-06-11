# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def bstToGst(self, root: TreeNode) -> TreeNode:
        if root is None:
            return None
        
        currSum = [0]
        self.recurse(root, currSum)
        return root
        
    def recurse(self, root, currSum):
        if root is None:
            return
        self.recurse(root.right, currSum)
        root.val = root.val + currSum[0]
        currSum[0] = root.val
        self.recurse(root.left, currSum)