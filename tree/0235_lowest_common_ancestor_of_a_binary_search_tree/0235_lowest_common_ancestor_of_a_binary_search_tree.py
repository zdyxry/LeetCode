class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def lowestCommonAncestor(self, root, p, q):
        s, b = sorted([p.val, q.val])
        while not s <= root.val <= b:
            root = root.left if s <= root.val else root.right
        return root


root = TreeNode(6)
root.left = TreeNode(2)
root.right = TreeNode(8)
res = Solution().lowestCommonAncestor(root,TreeNode(2), TreeNode(8))
print(res.val)