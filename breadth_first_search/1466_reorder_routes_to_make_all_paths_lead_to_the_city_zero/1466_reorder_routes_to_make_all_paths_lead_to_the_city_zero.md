1466. Reorder Routes to Make All Paths Lead to the City Zero


Medium


There are n cities numbered from 0 to n-1 and n-1 roads such that there is only one way to travel between two different cities (this network form a tree). Last year, The ministry of transport decided to orient the roads in one direction because they are too narrow.

Roads are represented by connections where connections[i] = [a, b] represents a road from city a to b.

This year, there will be a big event in the capital (city 0), and many people want to travel to this city.

Your task consists of reorienting some roads such that each city can visit the city 0. Return the minimum number of edges changed.

It's guaranteed that each city can reach the city 0 after reorder.

 

Example 1:

![1](1466-1.png)

```
Input: n = 6, connections = [[0,1],[1,3],[2,3],[4,0],[4,5]]
Output: 3
Explanation: Change the direction of edges show in red such that each node can reach the node 0 (capital).
```

Example 2:

![2](1466-2.png)

```
Input: n = 5, connections = [[1,0],[1,2],[3,2],[3,4]]
Output: 2
Explanation: Change the direction of edges show in red such that each node can reach the node 0 (capital).
```

Example 3:

```
Input: n = 3, connections = [[1,0],[2,0]]
Output: 0
```
 

Constraints:

2 <= n <= 5 * 10^4  
connections.length == n-1  
connections[i].length == 2  
0 <= connections[i][0], connections[i][1] <= n-1  
connections[i][0] != connections[i][1]


## 方法


```go
func minReorder(n int, connections [][]int) int {
    from, to := 1, 2
    m := map[int]map[int]int{}
    for i := 0; i < n; i++ {
        m[i] = make(map[int]int)
    }
    for _, c := range connections {
        m[c[0]][c[1]] = from
        m[c[1]][c[0]] = to
    }
    
    fmt.Println(m)

    modify := 0
    stack := []int{0}
    visited := map[int]bool{0: true}
    for i := 0; i < len(stack); i++ {
        for k, v := range m[stack[i]] {
            if visited[k] {
                continue
            }
            if v == from {
                modify++
            }
            stack = append(stack, k)
            visited[k] = true
        }
    }
    return modify
}
```


```python
class Solution:
    def minReorder(self, n: int, connections: List[List[int]]) -> int:
        edges = { (a,b) for a, b in connections } # instantly check if a->b
        neighbors = { i:[] for i in range(n)}     # adjacent cities
        visit = set()                             # visit each city once
        changes = 0
        
        for a, b in connections:
            neighbors[a].append(b)
            neighbors[b].append(a)
        
        # dfs from City 0, count outgoing edges
        def dfs(city):
            nonlocal edges, neighbors, visit, changes
            
            for neighbor in neighbors[city]:
                if neighbor in visit:
                    continue
                # neighbor can't reach city
                if (neighbor, city) not in edges:
                    changes += 1
                visit.add(neighbor)
                dfs(neighbor)
        
        visit.add(0)
        dfs(0)
        return changes
```