
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def balanceBST(self, root):
        def dfs(node):
            if not node:
                return
            dfs(node.left)
            values.append(node.val)
            dfs(node.right)

        values = []
        dfs(root)

        if not values:
            return None

        def tree(x):
            if not x:
                return None
            mid = len(x) // 2
            ans = TreeNode(x[mid])
            ans.left = tree(x[:mid])
            ans.right = tree(x[mid+1:])
            return ans
        
        return tree(values)

            