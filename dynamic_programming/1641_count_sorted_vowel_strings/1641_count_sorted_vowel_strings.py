class Solution:
    def countVowelStrings(self, n: int) -> int:
        dp=[1,1,1,1,1]
        if n==1:
            return 5

        for i in range(2,n+1):
            a = dp[0]
            e = dp[1]+dp[0]
            i = dp[0]+dp[1]+dp[2]
            o = dp[0]+dp[1]+dp[2]+dp[3]
            u = sum(dp)

            dp[0]=a
            dp[1]=e
            dp[2]=i
            dp[3]=o
            dp[4]=u
        
        return sum(dp)

n = 1
res = Solution().countVowelStrings(n)
print(res)