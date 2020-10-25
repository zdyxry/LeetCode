from typing import List

class Solution:
    def diStringMatch(self, S: str) -> List[int]:
        lo, hi = 0, len(S)
        ans = []
        for x in S:
            if x == 'I':
                ans.append(lo)
                lo += 1
            else:
                ans.append(hi)
                hi -= 1

        return ans + [lo]


S = "IDID"
res = Solution().diStringMatch(S)
print(res)