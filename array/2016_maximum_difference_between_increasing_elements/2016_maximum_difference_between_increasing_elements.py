class Solution:
    def maximumDifference(self, nums: List[int]) -> int:
        pre_min = nums[0]
        result = -1
        for i in nums[1:]:
            result = max(result, i - pre_min)
            if i < pre_min:
                pre_min = min(pre_min, i )
        return result if result else -1