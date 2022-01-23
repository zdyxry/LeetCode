1971. Find if Path Exists in Graph


Easy

There is a bi-directional graph with n vertices, where each vertex is labeled from 0 to n - 1 (inclusive). The edges in the graph are represented as a 2D integer array edges, where each edges[i] = [ui, vi] denotes a bi-directional edge between vertex ui and vertex vi. Every vertex pair is connected by at most one edge, and no vertex has an edge to itself.

You want to determine if there is a valid path that exists from vertex source to vertex destination.

Given edges and the integers n, source, and destination, return true if there is a valid path from source to destination, or false otherwise.

 

Example 1:

```
Input: n = 3, edges = [[0,1],[1,2],[2,0]], source = 0, destination = 2
Output: true
Explanation: There are two paths from vertex 0 to vertex 2:
- 0 → 1 → 2
- 0 → 2
```

Example 2:


```
Input: n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], source = 0, destination = 5
Output: false
Explanation: There is no path from vertex 0 to vertex 5.
```

Constraints:

```
1 <= n <= 2 * 105
0 <= edges.length <= 2 * 105
edges[i].length == 2
0 <= ui, vi <= n - 1
ui != vi
0 <= source, destination <= n - 1
There are no duplicate edges.
There are no self edges.
```

## 方法



```go
func validPath(n int, edges [][]int, start int, end int) bool {
    graph := make([][]int, n)
    for _, e := range edges {
        graph[e[0]] = append(graph[e[0]], e[1])
        graph[e[1]] = append(graph[e[1]], e[0])
    }
    visited := make(map[int]bool)
    q := []int{ start }
    visited[start] = true
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        if cur == end { return true }
        for _, node := range graph[cur] {
            if !visited[node] {
                visited[node] = true
                q = append(q, node)
            }
        }
    }
    return false
}
```


```python
class Solution:
    def validPath(self, n: int, edges: List[List[int]], source: int, destination: int) -> bool:
        d = {}
        for i in edges:
            a, b  = i 
            d.setdefault(a, [])
            d.setdefault(b, [])
            d[a].append(b)
            d[b].append(a)
        
        queue = [source]
        vis = set()
        vis.add(source)
        while queue:
            h = queue.pop()
            if h == destination:
                return True
            for i in d[h]:
                if i not in vis:
                    queue.append(i)
                    vis.add(i)
        return False
```