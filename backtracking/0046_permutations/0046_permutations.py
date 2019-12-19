
class Solution(object):
    def permute(self, nums):
        if len(nums) == 0:
            return
        result = []
        buffer = []
        def permutation(nums, buffer):
            if len(nums) == 0:
                result.append([i for i in buffer])
                return
            for i, ch in enumerate(nums):
                buffer.append(ch)
                permutation(nums[:i] + nums[i+1:], buffer)
                buffer.pop()
        permutation(nums, buffer)
        return result


nums = [1,2,3]
res = Solution().permute(nums)
print(res)