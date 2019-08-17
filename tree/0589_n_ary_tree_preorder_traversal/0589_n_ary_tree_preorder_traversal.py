class Node(object):
    def __init__(self, val, children):
        self.val = val
        self.children = children


class Solution(object):
    def preorder(self, root):
        result = []
        def helper(node):
            if node is None:
                return
            result.append(node.val)
            for child in node.children:
                helper(child)
            
        helper(root)
        return result

    def preorder2(self, root):
        if not root:
            return []
        
        traversal = []
        stack = [root]
        while stack:
            cur = stack.pop()
            traversal.append(cur.val)
            stack.extend(reversed(cur.children))
        return traversal