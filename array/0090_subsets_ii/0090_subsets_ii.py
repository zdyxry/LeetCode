class Solution(object):
    def subsetsWithDup(self, nums):
        nums.sort()
        result = [[]]
        previous_size = 0

        for i in xrange(len(nums)):
            print("i is", i)
            size = len(result)
            for j in xrange(size):
                print("j is", j)
                if i == 0 or nums[i] != nums[i-1] or j >= previous_size:
                    result.append(list(result[j]))
                    result[-1].append(nums[i])
                    print(result)
            previous_size = size
        return result

nums = [1,2,2]
res = Solution().subsetsWithDup(nums)
print(res)