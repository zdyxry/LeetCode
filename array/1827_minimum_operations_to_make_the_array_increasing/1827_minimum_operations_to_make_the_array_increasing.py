from typing import List

class Solution:
    def minOperations(self, nums: List[int]) -> int:
        cnt = 0
        for i in range(1, len(nums)):
            if nums[i] <= nums[i-1]:
                cnt += nums[i-1] + 1 - nums[i]
                nums[i] = nums[i-1] + 1
        return cnt
                


nums = [1,1,1]
res = Solution().minOperations(nums)
print(res)