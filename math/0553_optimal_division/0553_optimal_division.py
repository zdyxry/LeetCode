from typing import List

class Solution:
    def optimalDivision(self, nums: List[int]) -> str:
        n = len(nums)
        if n == 0: return ""
        if n == 1: return str(nums[0])
        if n == 2: return str(nums[0]) + "/" + str(nums[1])
        res = str(nums[0])
        res += "/(" + str(nums[1])
        for i in range(2, n):
            res += "/" + str(nums[i])
        res += ")"
        return res



nums = [1000,100,10,2]
res = Solution().optimalDivision(nums)
print(res)