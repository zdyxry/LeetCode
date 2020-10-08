

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class BSTIterator:

    def __init__(self, root: TreeNode):
        self.index = 0
        self.res = []
        self.dfs(root)

    def next(self) -> int:
        """
        @return the next smallest number
        """
        tmp = self.index
        self.index += 1
        return self.res[tmp]

    def hasNext(self) -> bool:
        """
        @return whether we have a next smallest number
        """
        if self.index == len(self.res):
            return False
        return True

    def dfs(self, root):
        if not root:
            return
        self.dfs(root.left)
        self.res.append(root.val)
        self.dfs(root.right)
