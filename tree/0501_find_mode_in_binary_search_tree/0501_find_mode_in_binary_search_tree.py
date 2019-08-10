# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def findMode(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        if not root:
            return []
        
        self.count = 0
        self.mode_count = 0
        self.modes = []
        self.last_seen = None
        
        def helper(node):
            
            if node.left:
                helper(node.left)
                
            if node.val != self.last_seen:
                if self.count > self.mode_count:
                    self.mode_count = self.count
                    self.modes = [self.last_seen]
                elif self.count == self.mode_count:
                    self.modes.append(self.last_seen)
                self.last_seen = node.val
                self.count = 1
            else:
                self.count += 1
                
            if node.right:
                helper(node.right)
        
        helper(root)
        if self.count > self.mode_count:
            return [self.last_seen]
        elif self.count == self.mode_count:
            self.modes.append(self.last_seen)
            return self.modes
        else:
            return self.modes
        

root = TreeNode(1)
root.right = TreeNode(2)
root.right.left = TreeNode(2)
print(Solution().findMode(root))