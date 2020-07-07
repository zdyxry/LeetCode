939. Minimum Area Rectangle


Medium


Given a set of points in the xy-plane, determine the minimum area of a rectangle formed from these points, with sides parallel to the x and y axes.

If there isn't any rectangle, return 0.

 

Example 1:

```
Input: [[1,1],[1,3],[3,1],[3,3],[2,2]]
Output: 4
```

Example 2:

```
Input: [[1,1],[1,3],[3,1],[3,3],[4,1],[4,3]]
Output: 2
```

Note:

1 <= points.length <= 500  
0 <= points[i][0] <= 40000  
0 <= points[i][1] <= 40000  
All points are distinct.


## 方法


```go
func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}

func minAreaRect(points [][]int) int {
    set := map[[2]int]bool{}
    for _, point := range points {
        set[[2]int{point[0], point[1]}] = true
    }
    best := 0
    
    for i := 0; i < len(points); i++ {
        p1 := points[i]
        for j := i + 1; j < len(points); j++ {
            p2 := points[j]
            if p1[0] == p2[0] || p1[1] == p2[1] {
                continue
            }
            p3 := [2]int{p1[0], p2[1]}
            p4 := [2]int{p2[0], p1[1]}
            if !set[p3] || !set[p4] {
                continue
            }
            area := abs(p1[0]-p2[0]) * abs(p1[1] - p2[1])
            if best <= 0 || area < best {
                best = area
            }
        }
    }
    return best
}
```


```python
class Solution:
    def minAreaRect(self, points: List[List[int]]) -> int:
        area = 0
        mem = set()
        for x1,y1 in points:
            for x2,y2 in mem:
                if (x2,y1) in mem and (x1,y2) in mem:
                    if not area:
                        area = abs(x2-x1)*abs(y2-y1)
                    else:
                        if abs(x2-x1)*abs(y2-y1) < area:
                            area = abs(x2-x1)*abs(y2-y1)
            mem.add((x1,y1))
        return area
```