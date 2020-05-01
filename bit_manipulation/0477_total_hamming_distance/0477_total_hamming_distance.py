from typing import List


class Solution(object):
    def totalHammingDistance(self, nums: List[int]) -> int:
        res, n = 0, len(nums)
        for i in range(32):
            cnt_1 = 0
            for j in range(n):
                cnt_1 += (nums[j]>>1) & 1
            res += (n-cnt_1) * cnt_1
        return res


nums = [4,14,2]
res = Solution().totalHammingDistance(nums)
print(res)