class Node(object):
    def __init__(self, val, children):
        self.val = val
        self.children = children


class Solution(object):
    def postorder(self,root):
        res = []
        if root == None: 
            return res 
        
        stack =[root]
        while stack:
            curr = stack.pop()
            res.append(curr.val)
            stack.extend(curr.children)
        
        return reversed(res)

    def postorder2(self, root):
        res = []
        if root == None:
            return res 

        def recursion(root, res):
            for child in root.children:
                recursion(child, res)
            res.append(root.val)
        
        recursion(root, res)
        return res