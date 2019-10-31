class Solution(object):
    def findMaxLength(self, nums):
        result, count = 0, 0
        lookup = {0: -1}
        for i, num in enumerate(nums):
            count += 1 if num == 1 else -1
            if count in lookup:
                result = max(result, i - lookup[count])
            else:
                lookup[count] = i 
        
        return result

nums = [0,1]

res = Solution().findMaxLength(nums)
print(res)