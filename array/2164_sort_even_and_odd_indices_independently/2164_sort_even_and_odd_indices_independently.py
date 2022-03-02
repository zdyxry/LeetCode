class Solution:
    def sortEvenOdd(self, nums: List[int]) -> List[int]:
        even = sorted(nums[::2])
        odd = sorted(nums[1::2])[::-1]
        nums[::2] = even
        nums[1::2] = odd
        return nums
