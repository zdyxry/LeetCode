class Solution(object):
    def maxProduct(self, A):
        global_max, local_max, local_min = float('-inf'), 1, 1
        for x in A:
            local_max, local_min = max(x, local_max * x, local_min * x), min(x, local_max * x, local_min * x)
            global_max = max(global_max, local_max)
        return global_max

A = [2,3,-2,4]

res = Solution().maxProduct(A)
print(res)