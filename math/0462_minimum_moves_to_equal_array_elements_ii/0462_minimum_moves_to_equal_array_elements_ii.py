from typing import List

class Solution:
    def minMoves2(self, nums: List[int]) -> int:
        nums.sort()
        i, j = 0, len(nums)-1
        ans = 0
        while i<j:
            ans += nums[j]-nums[i]
            i += 1
            j -= 1     
        return ans


nums = [1,2,3]
res = Solution().minMoves2(nums)
print(res)