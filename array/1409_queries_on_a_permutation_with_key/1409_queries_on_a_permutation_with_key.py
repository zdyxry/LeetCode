import collections

from typing import List



class Solution:
    def processQueries(self, queries: List[int], m: int) -> List[int]:
        tree = [0] * ((2*m) + 1)
        res = []
        
        def update(i,val):
            while i<len(tree):
                tree[i]+=val
                i+=(i&(-i))
    
        def prefixSum(i):
            s=0
            while i>0:
                s+=tree[i]
                i-=(i&(-i))
            return s
        
        hmap = collections.defaultdict(int)
        
        for i in range(1,m+1):
            hmap[i] = i+m
            update(i+m,1)

        for i in queries:
            res.append(prefixSum(hmap[i])-1)
            update(hmap[i],-1)
            update(m,1)
            hmap[i] = m
            m-=1

        return res


queries = [3,1,2,1]
m = 5
res = Solution().processQueries(queries, m)
print(res)