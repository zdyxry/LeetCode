from typing import List

class Solution:
    def selfDividingNumbers(self, left: int, right: int) -> List[int]:
        res = []
        for i in range(left, right+1):
            for j in list(str(i)):
                if int(j) == 0:
                    break
                if i % int(j) != 0:
                    break
            else:
                res.append(i)
        return res

left = 1
right = 22
res = Solution().selfDividingNumbers(left, right)
print(res)