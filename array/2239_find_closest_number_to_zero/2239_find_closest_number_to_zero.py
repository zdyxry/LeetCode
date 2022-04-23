class Solution:
    def findClosestNumber(self, nums: List[int]) -> int:
        result = float('inf')
        min_value = float('inf')
        for n in nums:
            if abs(n-0) == min_value:
                result = max(result, n)
            if abs(n-0) < min_value:
                result = n
                min_value = abs(n-0)
        return result
        