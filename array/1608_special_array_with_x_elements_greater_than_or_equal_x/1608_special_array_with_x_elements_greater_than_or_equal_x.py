from typing import List

class Solution:
    def specialArray(self, nums: List[int]) -> int:
        nums.sort(reverse=True)
        left, right = 0, len(nums)
        while left < right:
            mid = left + (right - left) // 2
            if mid < nums[mid]:
                left = mid + 1
            else:
                right = mid
        return -1 if left < len(nums) and left == nums[left] else left


nums = [3,5]
res = Solution().specialArray(nums)
print(res)