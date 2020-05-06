from typing import List


class Solution:
    def kLengthApart(self, nums: List[int], k: int) -> bool:
        zero_num = 1e5
        for x in nums:
            if x == 1:
                if zero_num < k:
                    return False
                zero_num = 0
            else:
                zero_num += 1
        return True


nums = [1,0,0,0,1,0,0,1]
k = 2
res = Solution().kLengthApart(nums, k)
print(res)