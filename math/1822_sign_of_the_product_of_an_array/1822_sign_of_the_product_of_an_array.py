from typing import List

class Solution:
    def arraySign(self, nums: List[int]) -> int:
        ans=1
        for i in nums:
            if i==0:
                return 0
            elif i<0:
                ans*=-1
        return ans


nums= [-1,-2,-3,-4,3,2,1]
res = Solution().arraySign(nums)
print(res)