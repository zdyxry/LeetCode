import collections

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def countPairs(self, root: TreeNode, distance: int) -> int:
        if not root: return 0
        self.res = 0
        def dfs(node, depth):
            if not node.left and not node.right: return collections.Counter([depth])
            l, r = collections.Counter(), collections.Counter()
            if node.left:
                l = dfs(node.left, depth + 1)
            if node.right:
                r = dfs(node.right, depth + 1)
            for k in l:
                for j in r:
                    if j + k <= distance +2 * depth:
                        self.res += l[k] * r[j]
            return l + r
        dfs(root, 0)
        return self.res
