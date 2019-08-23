class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def searchBST(self, root, val):
        if not root:
            return None
        if root.val == val:
            return root
        if root.val > val:
            return self.searchBST(root.left, val)
        if root.val < val:
            return self.searchBST(root.right, val)

root = TreeNode(2)
root.left = TreeNode(1)
root.right = TreeNode(3)
res = Solution().searchBST(root, 3)
print(res.val)