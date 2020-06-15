from typing import List

class Solution:
    def runningSum(self, nums: List[int]) -> List[int]:
        res = []
        sum = 0
        for i in nums:
            sum += i
            res.append(sum)
        return res


nums = [1,2,3,4]
res = Solution().runningSum(nums)
print(res)