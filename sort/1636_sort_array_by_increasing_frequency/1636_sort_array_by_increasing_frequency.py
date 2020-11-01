from typing import List

class Solution:
    def frequencySort(self, nums: List[int]) -> List[int]:
        return sorted(nums, key = lambda n: (nums.count(n), -n) )


nums = [1,1,2,2,2,3]
res = Solution().frequencySort(nums)
print(res)