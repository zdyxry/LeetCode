
class Solution:
    def findLexSmallestString(self, s: str, a: int, b: int) -> str:
        b = b%len(s)
        def addOpt(s):
            ans = [c for c in s]
            for i in range(1,len(s),2):               
                ans[i]=str((int(s[i])+a))[-1]
            return ''.join(ans)
        def lunOpt(s):
            return s[-b:]+s[:-b]
        alls = {s}  
        def dfs(s):           
            alls.add(s)          
            s1 = addOpt(s)
            s2 = lunOpt(s)          
            if s1 not in alls:                
                dfs(s1)           
            if s2 not in alls:
                dfs(s2)
        dfs(s)
        return min(alls)



s = "5525"
a = 9
b = 2
res = Solution().findLexSmallestString(s, a, b)
print(res)