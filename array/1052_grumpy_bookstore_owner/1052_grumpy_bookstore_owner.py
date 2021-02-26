from typing import List

class Solution:
    def maxSatisfied(self, customers: List[int], grumpy: List[int], X: int) -> int:
        res = 0
        for i in range(len(customers)):
            if grumpy[i] == 0:
                res += customers[i]
                customers[i] = 0
        max_count = 0
        val = 0
        for i in range(len(customers)):
            if i < X:
                val += customers[i]
                continue
            max_count = max(val, max_count)
            val -= customers[i-X]
            val += customers[i]
        max_count = max(val, max_count)
        return res + max_count



customers = [1,0,1,2,1,1,7,5]
grumpy = [0,1,0,1,0,1,0,1]
X = 3
res = Solution().maxSatisfied(customers, grumpy, X)
print(res)