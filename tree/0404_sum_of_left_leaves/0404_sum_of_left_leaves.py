class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def sumOfLeftLeaves(self, root):
        def sumOfLeftLeavesHelper(root, is_left):
            if not root:
                return 0
            if not root.left and not root.right :
                return root.val if is_left else 0
            return sumOfLeftLeavesHelper(root.left, True) + sumOfLeftLeavesHelper(root.right, False)

        return sumOfLeftLeavesHelper(root, False)


root = TreeNode(1)
root.left = TreeNode(2)
root.right = TreeNode(3)
res = Solution().sumOfLeftLeaves(root)
print(res)