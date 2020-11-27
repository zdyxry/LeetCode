from typing import List


class Solution:
    def waysToMakeFair(self, nums: List[int]) -> int:
        n = len(nums)
        odd_prefix, even_prefix = [0] * (n + 1), [0] * (n + 1)
        for i, num in enumerate(nums):
            odd_prefix[i] += odd_prefix[i-1]
            even_prefix[i] += even_prefix[i-1]
            if i % 2 == 0:
                even_prefix[i] += num
            else:
                odd_prefix[i] += num
        result = 0
        for i, num in enumerate(nums):
            even_after_removed = even_prefix[i-1] + odd_prefix[n-1] - odd_prefix[i]
            odd_after_removed = odd_prefix[i-1] + even_prefix[n-1] - even_prefix[i]
            result += (even_after_removed == odd_after_removed)
        return result


nums = [2,1,6,4]
res = Solution().waysToMakeFair(nums)
print(res)