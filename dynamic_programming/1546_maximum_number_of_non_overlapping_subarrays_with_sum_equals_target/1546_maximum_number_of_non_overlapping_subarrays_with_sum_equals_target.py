from typing import List


class Solution:
    def maxNonOverlapping(self, nums: List[int], target: int) -> int:
        s = {0}
        cur_sum = 0
        ans = 0
        for num in nums:
            cur_sum += num
            if cur_sum - target in s:
                s = {0}
                cur_sum = 0
                ans += 1
            else:
                s.add(cur_sum)
        return ans



nums = [1,1,1,1,1]
target = 2
res = Solution().maxNonOverlapping(nums, target)
print(res)