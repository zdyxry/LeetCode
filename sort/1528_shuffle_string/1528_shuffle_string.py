from typing import List

class Solution:
    def restoreString(self, s: str, indices: List[int]) -> str:
        _s = ['0'] * len(s)
        for idx, val in enumerate(s):
            _s[indices[idx]] = val
        return ''.join(_s)


s = "codeleet"
indices = [4,5,6,7,0,2,1,3]
res = Solution().restoreString(s, indices)
print(res)