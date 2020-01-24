
class Solution(object):
    def largestNumber(self, nums):
        num = [str(x) for x in nums]
        num.sort(cmp=lambda x, y: cmp(y + x, x + y))
        largest = "".join(num)
        return largest.lstrip("0") or "0"

nums = [10, 2]
res = Solution().largestNumber(nums)
print(res)