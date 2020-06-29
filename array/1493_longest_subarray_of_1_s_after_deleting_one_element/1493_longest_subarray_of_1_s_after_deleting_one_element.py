from typing import List

class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        cnt1, cnt2 = 0, 0
        maxlen = 0
        for num in nums:
            if num == 1:
                cnt1 += 1
                cnt2 += 1
                maxlen = max(maxlen, cnt1)
            else:
                cnt1 = cnt2
                cnt2 = 0
        if cnt1 == len(nums):
            return len(nums)-1
        return maxlen


nums = [1,1,0,0,1]
res = Solution().longestSubarray(nums)
print(res)