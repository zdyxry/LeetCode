from typing import List

class Solution:
    def getMaxLen(self, nums: List[int]) -> int:
        pre = -1
        l = []
        res = 0
        for i, num in enumerate(nums):
            if num < 0 :
                l.append(i)
            elif num == 0:
                l = []
                pre = i
            if len(l) % 2 == 0:
                res = max(res, i - pre)
            else:
                res = max(res, i - l[0])
        return res


nums = [1,-2,-3,4]
res = Solution().getMaxLen(nums)
print(res)