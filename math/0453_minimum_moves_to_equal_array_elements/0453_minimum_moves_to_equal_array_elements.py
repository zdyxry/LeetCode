from typing import List

class Solution:
    def minMoves(self, nums: List[int]) -> int:
        sum = 0
        minmum = min(nums)
        for i in nums:
            sum += i-minmum
        return sum


nums = [1,2,3]
res = Solution().minMoves(nums)
print(res)