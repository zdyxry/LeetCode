from typing import List

class Solution:
    def check(self, nums: List[int]) -> bool:
        B = sorted(nums)
        for i in range(len(nums)):
            if nums[i:]+nums[:i] == B:
                return True
        else:
            return False


nums = [3,4,5,1,2]
res = Solution().check(nums)
print(res)