
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def getAllElements(self, root1: TreeNode, root2: TreeNode) -> List[int]:
        ans = list()

        def dfs(node):
            if not node:
                return
            ans.append(node.val)
            dfs(node.left)
            dfs(node.right)
        
        dfs(root1)
        dfs(root2)
        ans.sort()
        return ans
