class Solution:
    def intersection(self, nums: List[List[int]]) -> List[int]:
        result = nums[0]
        for arr in nums[1:]:
            result = [i for i in result if i in arr]
        result.sort()
        return result

            