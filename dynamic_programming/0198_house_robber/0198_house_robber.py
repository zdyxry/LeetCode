
class Solution(object):
    def rob(self, nums):
        if nums == []:
            return 0
        if len(nums) == 1 :
            return nums[0]
        
        best = [nums[0], max(nums[0], nums[1])]
        for i, num in enumerate(nums[2:]):
            best.append(max(best[-2]+num, best[-1]))
        print(best)
        return best[-1]


nums = [1,2,3,1]
res = Solution().rob(nums)
print(res)