1443. Minimum Time to Collect All Apples in a Tree


Medium


Given an undirected tree consisting of n vertices numbered from 0 to n-1, which has some apples in their vertices. You spend 1 second to walk over one edge of the tree. Return the minimum time in seconds you have to spend in order to collect all apples in the tree starting at vertex 0 and coming back to this vertex.

The edges of the undirected tree are given in the array edges, where edges[i] = [fromi, toi] means that exists an edge connecting the vertices fromi and toi. Additionally, there is a boolean array hasApple, where hasApple[i] = true means that vertex i has an apple, otherwise, it does not have any apple.

 

Example 1:

![](1443-1.png)

```
Input: n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,false,true,false,true,true,false]
Output: 8 
Explanation: The figure above represents the given tree where red vertices have an apple. One optimal path to collect all apples is shown by the green arrows.  
```

Example 2:

![](1443-2.png)

```
Input: n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,false,true,false,false,true,false]
Output: 6
Explanation: The figure above represents the given tree where red vertices have an apple. One optimal path to collect all apples is shown by the green arrows.  
```

Example 3:

```
Input: n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,false,false,false,false,false,false]
Output: 0
```
 

Constraints:

1 <= n <= 10^5  
edges.length == n-1  
edges[i].length == 2  
0 <= fromi, toi <= n-1  
fromi < toi  
hasApple.length == n  


## 方法

```go
func minTime(n int, edges [][]int, hasApple []bool) int {
    graph := make(map[int][]int)
	for _, v := range edges {
		a := v[0]
		b := v[1]
		if _, ok := graph[a]; !ok {
			graph[a] = []int{}
		}
		graph[a] = append(graph[a], b)
		if _, ok := graph[b]; !ok {
			graph[b] = []int{}
		}
		graph[b] = append(graph[b], a)
	}
	set := make(map[int]bool)
	set[0] = true
    return dfs(0, graph, set, hasApple)
}

func dfs(cur int, graph map[int][]int, set map[int]bool, has []bool) int {
	set[cur] = true
	res:=0
	for _, v:=range graph[cur]{
		if !set[v]{
			res += dfs(v, graph, set, has) 	
		}
	}
	if cur==0{
		return res
	}
	if res>0||has[cur]{
		res+=2
	}
	return res
}
```



```python
class Solution:
    def minTime(self, n: int, edges: List[List[int]],
                hasApple: List[bool]) -> int:
        # 初始化路径
        maps = collections.defaultdict(list)
        for e in edges:
            maps[e[0]].append(e[1])

        def dfs(i):
            selfOrChildHasApple = hasApple[i]
            for nex in maps[i]:
                selfOrChildHasApple |= dfs(nex)
            if not selfOrChildHasApple:
                # 从字典中移除自身或子树都没有苹果的节点
                del maps[i]
            return selfOrChildHasApple

        dfs(0)
        # 字典个数即为最终有效节点的个数
        # 但是有可能有效节点为0, 所以需要max一下
        return max(0, 2 * (len(maps) - 1))
```