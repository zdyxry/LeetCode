class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def inorderTraversal(self, root):
        res, stack = [], []
        while True:
            while root:
                stack.append(root)
                root = root.left
            if not stack:
                return res
            node = stack.pop()
            res.append(node.val)
            root = node.right

root = TreeNode(1)
root.right = TreeNode(2)
root.right.left = TreeNode(3)

res = Solution().inorderTraversal(root)
print(res)