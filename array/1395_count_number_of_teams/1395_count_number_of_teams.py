class Solution(object):
    def numTeams(self, rating):
        """
        :type rating: List[int]
        :rtype: int
        """
        n = len(rating)
        ans = 0
        for i in range(1,n-1):
            l1,r1= 0,0
            l2,r2 =0,0
            for j in range(i-1,-1,-1):
                if rating[j] < rating[i]:
                    l1 += 1
                else:
                    l2 += 1
            for j in range(i+1,n):
                if rating[j] > rating[i]:
                    r1 += 1
                else:
                    r2 += 1
            ans += l1*r1 + l2*r2
        return ans


rating = [2,5,3,4,1]
res = Solution().numTeams(rating)
print(res)