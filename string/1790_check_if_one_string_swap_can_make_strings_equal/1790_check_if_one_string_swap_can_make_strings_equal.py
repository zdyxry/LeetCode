from typing import List

class Solution:
    def areAlmostEqual(self, s1: str, s2: str) -> bool:
        test = []
        if s1 == s2:
            return True
        for i in range(len(s1)):
            if s1[i] != s2[i]:
                test.append(i)
            else:
                pass
        if len(test) != 2:
            return False
        else:
            return(s1[test[0]] == s2[test[1]] and s1[test[1]] == s2[test[0]]) 



s1 = "bank"
s2 = "kanb"
res = Solution().areAlmostEqual(s1, s2)
print(res)