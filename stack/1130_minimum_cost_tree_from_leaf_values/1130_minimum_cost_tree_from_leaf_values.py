from typing import List

class Solution:
    def mctFromLeafValues(self, arr: List[int]) -> int:
        res, n = 0, len(arr)
        stack = [float('inf')]
        for a in arr:
            while stack[-1] <= a:
                mid = stack.pop()
                res += mid * min(stack[-1], a)
            stack.append(a)
        while len(stack) > 2:
            res += stack.pop() * stack[-1]
        return res


arr = [6,2,4,8]
res = Solution().mctFromLeafValues(arr)
print(res)