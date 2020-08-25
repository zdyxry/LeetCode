from typing import List

class Solution:
    def minOperations(self, nums: List[int]) -> int:
        count = 0
        while any(nums):
            for i in range(len(nums)):
                if nums[i] % 2:
                    nums[i] -= 1
                    count += 1
            if any(nums):
                for i in range(len(nums)):
                    nums[i] //= 2
                count += 1
        return count


nums  = [1,5]
res = Solution().minOperations(nums)
print(res)