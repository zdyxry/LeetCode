from typing import List

class Solution:
    def findContentChildren(self, g: List[int], s: List[int]) -> int:
        i, j = 0, 0
        g.sort()
        s.sort()
        while i < len(g) and j < len(s):
            if g[i] <= s[j]:
                i += 1
            j += 1
        return i
                
            


g = [10,9,8,7]
s = [5,6,7,8]
res = Solution().findContentChildren(g, s)
print(res)