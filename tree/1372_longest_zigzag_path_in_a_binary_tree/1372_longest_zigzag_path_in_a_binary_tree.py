
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
    

class Solution(object):
    def longestZigZag(self, root):
        def dfs(root):
            if not root:
                return [-1, -1, -1]
            left, right = dfs(root.left), dfs(root.right)
            return [left[1]+1, right[0]+1, max(left[1]+1, right[0]+1, left[2], right[2])]
        
        return dfs(root)[-1]


    