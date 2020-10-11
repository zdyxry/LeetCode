

class Solution:
    def maxDepth(self, s: str) -> int:
        ans=cur=0
        for l in s:
            if l=='(':
                cur+=1
            elif l==')':
                cur-=1
            ans=max(ans,cur)
        return ans


s = "(1+(2*3)+((8)/4))+1"
res = Solution().maxDepth(s)
print(res)