from typing import List


class Solution:
    def average(self, salary: List[int]) -> float:
        salary = sorted(salary)[1:-1]
        return sum(salary)/len(salary)



salary = [4000,3000,1000,2000]
res = Solution().average(salary)
print(res)