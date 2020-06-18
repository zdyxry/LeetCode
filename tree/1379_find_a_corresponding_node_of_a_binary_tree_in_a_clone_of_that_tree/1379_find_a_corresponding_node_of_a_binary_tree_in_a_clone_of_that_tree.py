# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def getTargetCopy(self, original: TreeNode, cloned: TreeNode, target: TreeNode) -> TreeNode:
        def traverse(original, cloned, target):
            if not original or not cloned:
                return None
            
            if original == target:
                return cloned
            
            return traverse(original.left, cloned.left, target) or \
                traverse(original.right, cloned.right, target)
        return traverse(original, cloned, target)