from typing import List
class Solution:
    def unhappyFriends(self, n: int, preferences: List[List[int]], pairs: List[List[int]]) -> int:
        dd = {}
        
        for i,x in pairs:
            dd[i] = preferences[i][:preferences[i].index(x)]
            dd[x] = preferences[x][:preferences[x].index(i)]
        
        ans = 0
            
        for i in dd:
            for x in dd[i]:
                if i in dd[x]:
                    ans += 1
                    break
        
        return ans


n = 4
preferences = [[1,2,3],[3,2,0],[3,1,0],[1,2,0]]
pairs = [[0,1],[2,3]]
res = Solution().unhappyFriends(n, preferences, pairs)
print(res)