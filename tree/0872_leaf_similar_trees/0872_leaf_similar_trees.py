# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def leafSimilar(self, root1, root2):
        return self.iterative(root1, []) == self.iterative(root2, [])

    def iterative(self, root, s):
        if root is not None:
            stack = []
            stack.append(root)
            while stack:
                x = stack.pop(-1)
                if x.left is None and x.right is None:
                    s.append(x.val)
                    continue
                if x.right is not None:
                    stack.append(x.right)
                if x.left is not None:
                    stack.append(x.left)
        return s

root1 = TreeNode(1)
root1.left = TreeNode(2)
root1.right = TreeNode(3)

root2 = TreeNode(4)
root2.left = TreeNode(2)
root2.right = TreeNode(3)

res = Solution().leafSimilar(root1, root2)
print(res)