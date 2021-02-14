class Solution:
    def minOperations(self, s: str) -> int:
        n=len(s)
        ans1=ans2=0
        for i in range(n):
            if i%2==0:
                if s[i]!='0':
                    ans2+=1
                else:
                    ans1+=1
            else:
                if s[i]!='0':
                    ans1+=1
                else:
                    ans2+=1
        return min(ans1,ans2)



s = "0100"
res = Solution().minOperations(s)
print(res)