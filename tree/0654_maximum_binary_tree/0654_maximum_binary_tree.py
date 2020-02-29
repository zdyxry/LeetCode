class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def constructMaximumBinaryTree(self, nums):
        if len(nums) == 0:
            return None
        val = max(nums)
        val_ind = nums.index(val)
        root = TreeNode(val)
        root.left = self.constructMaximumBinaryTree(nums[:val_ind])
        root.right = self.constructMaximumBinaryTree(nums[val_ind+1:])
        return root