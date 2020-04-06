from typing import List

class Solution:
    def minSubsequence(self, nums: List[int]) -> List[int]:
        N = len(nums)

        total = sum(nums)

        nums.sort()

        ret = []
        t = 0
        while True:
            x = nums.pop()

            t += x
            ret.append(x)
            if t * 2 > total: return ret


nums = [4,3,10,9,8]
res = Solution().minSubsequence(nums)
print(res)