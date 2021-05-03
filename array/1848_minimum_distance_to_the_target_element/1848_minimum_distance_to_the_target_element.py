from typing import List


class Solution:
    def getMinDistance(self, nums: List[int], target: int, start: int) -> int:
        i = j = start
        n = len(nums)
        while i>=0 or j<n:
            if i>=0 and nums[i]==target:
                return start-i
                              
            if j<n and nums[j]==target:
                return j-start
            i -= 1
            j += 1
        return 0



nums = [1,2,3,4,5]
target = 5
start = 3
res = Solution().getMinDistance(nums, target, start)
print(res)