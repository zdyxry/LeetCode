class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def pathSum(self, root, sum):
        def dfs(root, prevSum, sum):
            if not root:
                return
            currSum = prevSum + root.val
            if currSum - sum in rec:
                self.count += rec[currSum - sum]
            if currSum in rec:
                rec[currSum] += 1
            else:
                rec[currSum] =1
            dfs(root.left, currSum, sum)
            dfs(root.right, currSum, sum)
            rec[currSum] -= 1

        self.count = 0
        rec = {0:1}
        dfs(root, 0, sum)
        return self.count

root = TreeNode(10)
root.left = TreeNode(5)
root.right = TreeNode(-3)
root.left.left = TreeNode(3)
root.left.right = TreeNode(2)
root.right.right = TreeNode(11)
res = Solution().pathSum(root, 8)
print(res)