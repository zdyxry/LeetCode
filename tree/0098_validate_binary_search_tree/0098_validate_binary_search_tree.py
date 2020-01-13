class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def isValidBST(self, root):
        def helper(node, lower = float('-inf'), upper = float('inf')):
            if not node:
                return True

            val = node.val
            if val <= lower or val >= upper:
                return False
            
            if not helper(node.right, val, upper):
                return False

            if not helper(node.left, lower, val):
                return False

            return True

        return helper(root)