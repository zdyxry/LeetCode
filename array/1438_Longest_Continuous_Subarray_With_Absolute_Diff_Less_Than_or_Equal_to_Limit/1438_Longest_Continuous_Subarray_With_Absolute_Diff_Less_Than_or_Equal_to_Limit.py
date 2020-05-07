from typing import List


class Solution:
    def longestSubarray(self, nums: List[int], limit: int) -> int:
        minimal, maximal = float('inf'), float('-inf')
        size, current_size_start_number = 0, None
        for i in range(len(nums)):
            maximal = max(maximal, nums[i])
            minimal = min(minimal, nums[i])
            if abs(maximal - minimal) <= limit:
                size += 1
            else:
                current_size_start_number = nums[i-size]
                if current_size_start_number == minimal:
                    minimal = min(nums[i-size+1:i+1])
                elif current_size_start_number == maximal:
                    maximal = max(nums[i-size+1:i+1])
        return size


nums = [8,2,4,7]
limit = 4
res = Solution().longestSubarray(nums, limit)
print(res)