from collections import Counter

class Solution:
    def halvesAreAlike(self, s: str) -> bool:
        a = Counter(s[:len(s)//2])
        b = Counter(s[len(s)//2:])
        A=B=0
        for i in ('a','e','i','o','u','A','E','I','O','U'):
            A+=a[i]
            B+=b[i]
        return A==B



s = "book"
res = Solution().halvesAreAlike(s)
print(res)