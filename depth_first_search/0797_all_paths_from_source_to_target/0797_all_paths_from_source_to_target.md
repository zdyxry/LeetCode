797. All Paths From Source to Target


Medium


Given a directed, acyclic graph of N nodes.  Find all possible paths from node 0 to node N-1, and return them in any order.

The graph is given as follows:  the nodes are 0, 1, ..., graph.length - 1.  graph[i] is a list of all nodes j for which the edge (i, j) exists.

Example:
```
Input: [[1,2], [3], [3], []] 
Output: [[0,1,3],[0,2,3]] 
Explanation: The graph looks like this:
0--->1
|    |
v    v
2--->3
There are two paths: 0 -> 1 -> 3 and 0 -> 2 -> 3.
```

Note:

The number of nodes in the graph will be in the range [2, 15].

You can print different paths in any order, but you should keep the order of nodes inside one path.


## 方法


```go
func allPathsSourceTarget(graph [][]int) [][]int {
    var ans [][]int
    helper(graph, []int{0}, &ans) 
    return ans
}

func helper(graph [][]int, path []int, ans *[][]int) {
    pos := path[len(path)-1]
    target := len(graph)-1
    for _, child := range graph[pos] {
        if child == target {
            cp := make([]int, len(path))
            copy(cp, path)
            cp = append(cp, child)
            *ans = append(*ans, cp)
        } else {
            helper(graph, append(path, child), ans)
        }
    }
    return
}
```


```python
class Solution:
    def allPathsSourceTarget(self, graph: List[List[int]]) -> List[List[int]]:
        res = []
        self.dfs(res, graph, 0, [0])
        return res
    
    
    def dfs(self, res, graph, curr, path):
        
        if curr == len(graph) - 1:
            res.append(path)
            return
            
        for edge in graph[curr]:
            self.dfs(res, graph, edge, path+[edge])
```