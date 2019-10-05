class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None 


class Solution(object):
    def buildTree(self, preorder, inorder):
        lookup = {}
        for i, num in enumerate(inorder):
            lookup[num] = i
        return self.buildTreeRecu(lookup, preorder, inorder, 0, 0, len(inorder))

    def buildTreeRecu(self, lookup, preorder, inorder, pre_start, in_start, in_end):
        if in_start == in_end:
            return None 
        node = TreeNode(preorder[pre_start])
        i = lookup[preorder[pre_start]]
        node.left = self.buildTreeRecu(lookup, preorder, inorder, pre_start+1, in_start, i)
        node.right = self.buildTreeRecu(lookup, preorder, inorder, pre_start +1 +i - in_start, i +1, in_end)
        return node


preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]

res = Solution().buildTree(preorder, inorder)
print(res.val)