from typing import List


class Solution:
    def countSubTrees(self, n: int, edges: List[List[int]], labels: str) -> List[int]:
        from collections import defaultdict
        from collections import Counter

        edge_map = defaultdict(list)
        for edge in edges:
            edge_map[edge[0]].append(edge[1])
            edge_map[edge[1]].append(edge[0])
            
        def _dfs(i):
            visited.add(i)
            # 字符数字典
            data = Counter({labels[i]: 1})
            for nxt in edge_map[i]:
                if nxt in visited: continue # 去重
                # 整合子树的字符数
                data += _dfs(nxt)
            # 设置当前节点的结果字符数
            ans[i] = data[labels[i]]
            return data

        visited = set()
        ans = [1] * n
        _dfs(0)
        return ans


n = 7
edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]]
labels = "abaedcd"

res = Solution().countSubTrees(n, edges, labels)
print(res)