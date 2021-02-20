from typing import List

class Solution:
    def maxAbsoluteSum(self, nums: List[int]) -> int:
        pre_sum = 0
        max_sum = 0
        min_sum = 0
        for n in nums:
            pre_sum += n
            max_sum = max(pre_sum, max_sum)
            min_sum = min(pre_sum, min_sum)
        return max_sum - min_sum



nums = [1,-3,2,3,-4]
res = Solution().maxAbsoluteSum(nums)
print(res)