from typing import List


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

        
class Solution:
    def bstFromPreorder(self, preorder: List[int]) -> TreeNode:
        if len(preorder) == 0:
            return None
        leftList = []
        rightList = []
        root = TreeNode(preorder[0])
        for num in preorder[1:]:
            if num < preorder[0]:
                leftList.append(num)
            elif num > preorder[0]:
                rightList.append(num)
        root.left = self.bstFromPreorder(leftList)
        root.right = self.bstFromPreorder(rightList)
        return root