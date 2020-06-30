1496. Path Crossing


Easy


Given a string path, where path[i] = 'N', 'S', 'E' or 'W', each representing moving one unit north, south, east, or west, respectively. You start at the origin (0, 0) on a 2D plane and walk on the path specified by path.

Return True if the path crosses itself at any point, that is, if at any time you are on a location you've previously visited. Return False otherwise.

 

Example 1:

![1](1496-1.png)

```
Input: path = "NES"
Output: false 
Explanation: Notice that the path doesn't cross any point more than once.
```

Example 2:

![2](1496-2.png)

```
Input: path = "NESWW"
Output: true
Explanation: Notice that the path visits the origin twice.
```

Constraints:

1 <= path.length <= 10^4  
path will only consist of characters in {'N', 'S', 'E', 'W}

## 方法

```go
func isPathCrossing(path string) bool {
    visited := make(map[string] bool)
    
    visited["00"] = true
    
    x := 0
    y := 0
    
    for _, val := range path {  
        switch val {
        case 'N':
            y++
        case 'S':
            y--
        case 'E':
            x++
        default:
            x--
        }  
        
        
        key := fmt.Sprintf("%d%d", x, y) 
        
        _, ok := visited[key] 
        
        if ok {
            return true
        }
        
        visited[key] = true
    }
    
    return false
}
```


```python
class Solution:
    def isPathCrossing(self, path: str) -> bool:
        x = y = 0
        set = {(0, 0)}
        
        for c in path:
            if c == 'N':
                y += 1
            elif c == 'S':
                y -= 1
            elif c == 'E':
                x += 1
            else:
                x -= 1
            
            if (x, y) in set:
                return True
            set.add((x, y))
        return False
```