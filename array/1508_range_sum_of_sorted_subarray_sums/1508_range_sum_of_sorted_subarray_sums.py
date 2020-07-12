from typing import List


class Solution:
    def rangeSum(self, nums: List[int], n: int, left: int, right: int) -> int:
        res = []
        for i,n in enumerate(nums):
            res.append(n)
            total = n
            for j in nums[i+1:]:
                total += j
                res.append(total)
        res.sort()
        r = 0
        for m in res[left-1:right]:
            r += m
        return r
            
        

nums = [1,2,3,4]
n = 4
left = 1
right = 5
res = Solution().rangeSum(nums, n, left, right)
print(res)