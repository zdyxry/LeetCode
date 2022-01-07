class Solution:
    def maxProductDifference(self, nums: List[int]) -> int:
        sNums = sorted(nums)
        return sNums[-1]*sNums[-2] - sNums[0]*sNums[1]