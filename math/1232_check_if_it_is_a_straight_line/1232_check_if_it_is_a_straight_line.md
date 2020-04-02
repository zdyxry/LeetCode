1232. Check If It Is a Straight Line


Easy


You are given an array coordinates, coordinates[i] = [x, y], where [x, y] represents the coordinate of a point. Check if these points make a straight line in the XY plane.



Example 1:

![1](1232-1.jpg)

```
Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
Output: true
```

Example 2:

![2](1232-2.jpg)

```
Input: coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
Output: false
```
 

Constraints:

2 <= coordinates.length <= 1000  
coordinates[i].length == 2  
-10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4  
coordinates contains no duplicate point.   


## 方法

```go
func checkStraightLine(coordinates [][]int) bool {
    slope:=float32(coordinates[1][1]-coordinates[0][1])/float32(coordinates[1][0]-coordinates[0][0])
    for i:=2; i<len(coordinates); i++{
        if float32(coordinates[i][1]-coordinates[i-1][1])/float32(coordinates[i][0]-coordinates[i-1][0]) != slope{
            return false
        }
    }
    return true
}
```


```python
class Solution:
    def checkStraightLine(self, coordinates: List[List[int]]) -> bool:
        same_slop = True
        last_slop = None
        intial_point = coordinates[0]
        for point in coordinates[1:]:
            try:
                slop = (point[1] - intial_point[1])/(point[0] - intial_point[0])
            except ZeroDivisionError: 
                slop = float(inf)
            if last_slop== None:last_slop = slop
            elif slop == last_slop:continue
            else: 
                same_slop = False
                break
                
        return same_slop
```