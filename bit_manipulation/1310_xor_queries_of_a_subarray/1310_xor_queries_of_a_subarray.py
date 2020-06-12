from typing import List

class Solution:
    def xorQueries(self, arr: List[int], queries: List[List[int]]) -> List[int]:
        pre = [0]
        for num in arr:
            pre.append(pre[-1] ^ num)
        print(pre)
        ans = list()
        for x, y in queries:
            ans.append(pre[x] ^ pre[y + 1])
        return ans




arr = [1,3,4,8]
queries = [[0,1],[1,2],[0,3],[3,3]]
res = Solution().xorQueries(arr, queries)
print(res)