class Solution:
    def canBeIncreasing(self, nums: List[int]) -> bool:
        for i in range(len(nums)):
            tmp_nums = nums[0:i] + nums[i+1:]
            for j in range(len(tmp_nums)-1):
                if tmp_nums[j] >= tmp_nums[j+1]:
                    break
            else:
                return True
        return False



nums = [1,2,10,5,7]
res = Solution().canBeIncreasing(nums)
print(res)