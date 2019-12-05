
class Solution(object):
    def nextGreaterElements(self, nums):
        stack = []
        result = [-1]*len(nums)
        for enum, n in enumerate(nums):
            while stack and stack[-1][0] < n:
                result[stack[-1][1]] = n
                stack.pop()
            stack.append((n, enum))
        print result, stack
        if stack:
            for enum, n in enumerate(nums):
                while stack and stack[-1][0] < n:
                    result[stack[-1][1]] = n
                    stack.pop()

        return result

nums = [1,2,1]
res = Solution().nextGreaterElements(nums)
print(res)