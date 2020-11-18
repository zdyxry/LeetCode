import collections



class Solution:
    def closeStrings(self, word1: str, word2: str) -> bool:
        c1 = collections.Counter(word1)
        c2 = collections.Counter(word2)
        
        l1 = list(c1.values())
        l2 = list(c2.values())
        
        s1 = set(word1)
        s2 = set(word2)
        
        l1.sort()
        l2.sort()
        
        return l1 == l2 and s1 == s2



word1 = "abc"
word2 = "bca"
res = Solution().closeStrings(word1, word2)
print(res)