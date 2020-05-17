from typing import List


class Solution:
    def allPathsSourceTarget(self, graph: List[List[int]]) -> List[List[int]]:
        res = []
        self.dfs(res, graph, 0, [0])
        return res

    def dfs(self, res, graph, curr, path):
        if curr == len(graph) -1:
            res.append(path)
            return
        
        for edge in graph[curr]:
            self.dfs(res, graph, edge, path+[edge])


graph = [[1,2],[3],[3],[]]
res = Solution().allPathsSourceTarget(graph)
print(res)