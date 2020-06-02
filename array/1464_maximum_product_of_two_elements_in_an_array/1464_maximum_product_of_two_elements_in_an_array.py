from typing import List

class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        nums=sorted(nums)
        return ((nums[-2]-1)*(nums[-1]-1))


nums = [3,4,5,2]
res = Solution().maxProduct(nums)
print(res)