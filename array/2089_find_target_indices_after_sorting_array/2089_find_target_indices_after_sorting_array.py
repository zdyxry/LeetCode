class Solution:
    def targetIndices(self, nums: List[int], target: int) -> List[int]:
        result = []
        
        nums.sort()
        for i in range(len(nums)):
            if nums[i] == target:
                result.append(i)
            if nums[i] > target:
                break
        return result