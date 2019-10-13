class Solution(object):
    def majorityElement(self, nums):
        if not nums:
            return []

        num1 = -1
        num2 = -1
        count1 = 0
        count2 = 0

        for n in nums:
            if n == num1:
                count1 += 1
            elif n == num2:
                count2 += 1
            elif count1 == 0:
                num1 = n
                count1 = 1
            elif count2 == 0:
                num2 = n
                count2 = 1
            else:
                count1 -= 1
                count2 -= 1
        
        return [n for n in [num1, num2] if nums.count(n) > len(nums) // 3]

nums = [3,2,3]
res = Solution().majorityElement(nums)
print(res)