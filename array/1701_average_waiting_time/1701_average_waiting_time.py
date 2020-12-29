from typing import List

class Solution:
    def averageWaitingTime(self, customers: List[List[int]]) -> float:
        cur = customers[0][0]
        total = 0
        for c in customers:
            
            if cur >= c[0]:
                cur = cur + c[1]
                total += cur - c[0]
            else:
                cur = c[0] + c[1]
                total += c[1]
        return float(total) / len(customers)



customers = [[1,2],[2,5],[4,3]]
res =  Solution().averageWaitingTime(customers)
print(res)