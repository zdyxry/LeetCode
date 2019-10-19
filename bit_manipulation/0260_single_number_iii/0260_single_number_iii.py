class Solution(object):
    def singleNumber(self, nums):
        x_xor_y = 0
        for i in nums:
            x_xor_y ^= i
        
        bit = x_xor_y &~(x_xor_y - 1)

        x = 0
        for i in nums:
            if i & bit:
                x ^= i
            
        return [x, x^x_xor_y]


nums = [1,2,1,3,2,5]
res = Solution().singleNumber(nums)
print(res)