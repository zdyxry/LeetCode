from typing import List
from collections import Counter


class Solution:
    def numRabbits(self, answers: List[int]) -> int:
        cnt = Counter(answers)
        ans = 0
        
        for k, v in cnt.items():
            ans += (v+k)//(k+1)*(k+1)
        
        return ans


answers = [1,0,1,0,0]
res = Solution().numRabbits(answers)
print(res)