
class Solution(object):
    def createTargetArray(self, nums, index):
        target = []
        for n, i in zip(nums, index):
            target.insert(i, n)
        return target


nums = [0,1,2,3,4]
index = [0,1,2,2,1]
res = Solution().createTargetArray(nums, index)
print(res)