from typing import List

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class Solution:
    def allPossibleFBT(self, N: int) -> List[TreeNode]:
        res = []
        if N == 1:
            return [TreeNode(0)]
        if N % 2 == 0:
            return []
        
        left_num = 1
        right_num = N - 2
        
        while right_num > 0:
            left_tree = self.allPossibleFBT(left_num)
            right_tree = self.allPossibleFBT(right_num)
            for i in range(len(left_tree)):
                for j in range(len(right_tree)):
                    root = TreeNode(0)
                    root.left = left_tree[i]
                    root.right = right_tree[j]
                    res.append(root)
            left_num += 2
            right_num -= 2
        
        return res     
