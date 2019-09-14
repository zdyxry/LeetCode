import collections

class Solution(object):
    def fourSum(self, nums, target):
        nums, result, lookup = sorted(nums), [], collections.defaultdict(list)
        for i in xrange(0, len(nums) - 1):
            for j in xrange(i+1, len(nums)):
                lookup[nums[i]+nums[j]].append([i,j])

        for i in lookup.keys():
            if target - i in lookup:
                for x in lookup[i]:
                    for y in lookup[target-i]:
                        [a,b],[c,d] = x,y
                        if a is not c and a is not d and b is not c and b is not d:
                            quad = sorted([nums[a], nums[b], nums[c],nums[d]])
                            if quad not in result:
                                result.append(quad)
        return sorted(result)

nums = [1,0,-1,0,-2,2]
target = 0 

res = Solution().fourSum(nums, target)
print(res)