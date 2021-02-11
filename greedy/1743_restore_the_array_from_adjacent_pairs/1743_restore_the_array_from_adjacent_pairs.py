from typing import List
import collections



class Solution:
    def restoreArray(self, adjacentPairs: List[List[int]]) -> List[int]:
        dic = collections.defaultdict(list)
        n = len(adjacentPairs) + 1
        for x,y in adjacentPairs:
            dic[x].append(y)
            dic[y].append(x)
        head = [k for k,v in dic.items() if len(v)==1]
        res = [head[0]]
        while len(res) < n:
            i = res[-1]
            j = dic[i].pop()
            dic[j].remove(i)
            res.append(j)
        return res


adjacentPairs = [[2,1],[3,4],[3,2]]
res = Solution().restoreArray(adjacentPairs)
print(res)