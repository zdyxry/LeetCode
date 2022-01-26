class Solution:
    def minimumDifference(self, nums: List[int], k: int) -> int:
        nums.sort()
        left = 0
        right = k-1
        result = nums[-1]
        while right < len(nums):
            result = min(result, nums[right] - nums[left])
            left += 1
            right += 1
        return result

