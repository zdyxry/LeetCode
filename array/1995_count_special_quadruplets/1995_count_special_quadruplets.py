class Solution:
    def countQuadruplets(self, nums: List[int]) -> int:
        if len(nums) < 4:
            return 0
            
        result = 0
        for i in range(0, len(nums)):
            for j in range(i+1, len(nums)):
                for m in range(j+1, len(nums)):
                    for n in range(m+1, len(nums)):
                        if nums[i] + nums[j] + nums[m] == nums[n]:
                            result += 1
        return result
