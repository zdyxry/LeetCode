class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def convertBST(self, root):
        def covnertBSTHelper(root, cur_sum):
            if not root:
                return cur_sum
            if root.right:
                cur_sum = covnertBSTHelper(root.right, cur_sum)
            cur_sum += root.val
            root.val = cur_sum
            if root.left:
                cur_sum = covnertBSTHelper(root.left, cur_sum)
            return cur_sum
        covnertBSTHelper(root, 0)
        return root


root = TreeNode(5)
root.left = TreeNode(2)
root.right = TreeNode(13)
res = Solution().convertBST(root)
print(res.left.val)