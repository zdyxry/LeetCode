1552. Magnetic Force Between Two Balls


Medium


In universe Earth C-137, Rick discovered a special form of magnetic force between two balls if they are put in his new invented basket. Rick has n empty baskets, the ith basket is at position[i], Morty has m balls and needs to distribute the balls into the baskets such that the minimum magnetic force between any two balls is maximum.

Rick stated that magnetic force between two different balls at positions x and y is |x - y|.

Given the integer array position and the integer m. Return the required force.

 

Example 1:


```
Input: position = [1,2,3,4,7], m = 3
Output: 3
Explanation: Distributing the 3 balls into baskets 1, 4 and 7 will make the magnetic force between ball pairs [3, 3, 6]. The minimum magnetic force is 3. We cannot achieve a larger minimum magnetic force than 3.
```

Example 2:

```
Input: position = [5,4,3,2,1,1000000000], m = 2
Output: 999999999
Explanation: We can use baskets 1 and 1000000000.
```

Constraints:

n == position.length  
2 <= n <= 10^5  
1 <= position[i] <= 10^9  
All integers in position are distinct.  
2 <= m <= position.length


## 方法

```go
import (
    "fmt"
    "sort"
)

func search(dist, m int, position []int) bool{
    curr, res := 0, 1
    for i := 0; i < len(position) - 1; i ++ {
        curr += position[i+1] - position[i]
        if curr >= dist{
            res += 1
            curr = 0
        }
    }
    return res >= m
}

func maxDistance(position []int, m int) int {
    sort.Ints(position)
    n := len(position)
    l, r := 1, position[n-1] - position[0]
    
    for l < r {
        mid := (l + r + 1) / 2
        if search(mid, m, position) {
            l = mid
        } else {
            r = mid - 1
        }
    }
    return l
}
```



```python
class Solution:
    def maxDistance(self, position: List[int], m: int) -> int:
        position.sort()
        
        def check(x):
            cnt = 1
            t = position[0]
            for i in range(1, len(position)):
                if position[i]-t > x:
                    cnt += 1
                    t = position[i]
            return cnt >= m
        
        l, r = 0, position[-1]
        while l < r:
            mid = l + (r-l)//2
            if check(mid):
                l = mid+1
            else:
                r = mid
        return l
```