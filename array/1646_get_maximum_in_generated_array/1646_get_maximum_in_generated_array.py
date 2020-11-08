

class Solution:
    def getMaximumGenerated(self, n: int) -> int:
        if n == 0:
            return 0
        max_value = 1
        nums = [0] *(n + 1)
        nums[1] = 1
        for i in range(2, n+1):
            if i % 2 == 0:
                nums[i] = nums[i//2]
            else:
                nums[i] = nums[i//2] + nums[i//2 + 1]
            max_value = max(max_value, nums[i])
        return max_value


n = 7
res = Solution().getMaximumGenerated(n)
print(res)