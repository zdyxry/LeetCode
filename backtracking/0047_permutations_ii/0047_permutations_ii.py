
class Solution(object):
    def permuteUnique(self, nums):
        res = []
        nums.sort()
        self.dfs(nums, [], res)
        return res

    def dfs(self, remain, path, res):
        if not remain:
            res.append(path)
            return
        for i in range(len(remain)):
            if i > 0 and remain[i-1] == remain[i]:
                continue
            left = remain[:i] + remain[i+1:]
            self.dfs(left, path+[remain[i]], res)


nums = [1,1,2]
res = Solution().permuteUnique(nums)
print(res)