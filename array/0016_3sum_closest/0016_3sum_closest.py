class Solution(object):
    def threeSumColsest(self, nums, target):
        nums, result, min_diff, i = sorted(nums), float('inf'), float('inf'), 0
        while i < len(nums) - 2 :
            if i == 0 or nums[i] != nums[i-1]:
                j,k = i+1, len(nums) - 1
                while j < k:
                    diff = nums[i] + nums[j] + nums[k] - target 
                    if abs(diff) < min_diff:
                        min_diff = abs(diff)
                        result = nums[i] + nums[j] + nums[k]
                    if diff < 0:
                        j += 1
                    elif diff > 0 :
                        k -= 1
                    else:
                        return target 
            i += 1
        return result 

nums = [-1,2,1,-4]
target = 1

res = Solution().threeSumColsest(nums, target)
print(res)