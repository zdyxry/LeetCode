from typing import List


class Solution:
    def calPoints(self, ops: List[str]) -> int:
        res = []
        for op in ops:
            if op == 'C':
                res.pop()
            elif op == '+':
                res.append(res[-1] + res[-2])
            elif op == 'D':
                res.append(res[-1] * 2)
            else:
                res.append(int(op))
        return sum(res)


ops = ["5","2","C","D","+"]
res = Solution().calPoints(ops)
print(res)