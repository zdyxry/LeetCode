class Solution(object):
    def subarraySum(self, nums, k):
        count, p, d = 0, 0, {0:1}
        for i in nums:
            p += i
            if p-k in d:
                count += d[p-k]
            if p not in d:
                d[p] = 1
            else:
                d[p] += 1
        return count

nums = [1,1,1]
k = 2
res = Solution().subarraySum(nums, k)
print(res)