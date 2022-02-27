class Solution:
    def countElements(self, nums: List[int]) -> int:
        smallest = min(nums)
        largest = max(nums)
        res = 0
        for num in nums:
            if smallest < num < largest:
                res += 1
        return res
