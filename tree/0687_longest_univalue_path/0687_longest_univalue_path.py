class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def longestUnivaluePath(self, root):
        result = [0]
        def dfs(node):
            if not node:
                return 0
            left, right = dfs(node.left), dfs(node.right)
            if node.left and node.left.val == node.val:
                left = left + 1
            else:
                left = 0
            if node.right and node.right.val == node.val:
                right = right + 1
            else:
                right = 0
            result[0] = max(result[0], left + right)
            return max(left, right)
        
        dfs(root)
        return result[0]

root = TreeNode(1)
root.left = TreeNode(1)
root.right = TreeNode(1)
res = Solution().longestUnivaluePath(root)
print(res)