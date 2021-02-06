import collections
from typing import List

class Solution:
    def sumOfUnique(self, nums: List[int]) -> int:
        import collections
        return sum([i for i,v in collections.Counter(nums).items() if v <= 1])


nums = [1,2,3,2]
res = Solution().sumOfUnique(nums)
print(res)