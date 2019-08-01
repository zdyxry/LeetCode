class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def isBalanced(self, root):
        def getDepth(node):
            if not node:
                return 0
            left = getDepth(node.left)
            right = getDepth(node.right)

            if abs(left - right) > 1:
                self.res = False
            return 1 + max(left, right)

        self.res = True
        getDepth(root)
        return self.res


root = TreeNode(1)
root.left = TreeNode(2)
root.right = TreeNode(3)
root.right.right = TreeNode(4)
res = Solution().isBalanced(root)
print(res)