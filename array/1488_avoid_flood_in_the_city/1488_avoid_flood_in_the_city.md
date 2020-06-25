1488. Avoid Flood in The City


Medium


Your country has an infinite number of lakes. Initially, all the lakes are empty, but when it rains over the nth lake, the nth lake becomes full of water. If it rains over a lake which is full of water, there will be a flood. Your goal is to avoid the flood in any lake.

Given an integer array rains where:

```
rains[i] > 0 means there will be rains over the rains[i] lake.
rains[i] == 0 means there are no rains this day and you can choose one lake this day and dry it.
```

Return an array ans where:

```
ans.length == rains.length
ans[i] == -1 if rains[i] > 0.
ans[i] is the lake you choose to dry in the ith day if rains[i] == 0.
If there are multiple valid answers return any of them. If it is impossible to avoid flood return an empty array.
```

Notice that if you chose to dry a full lake, it becomes empty, but if you chose to dry an empty lake, nothing changes. (see example 4)

 

Example 1:

```
Input: rains = [1,2,3,4]
Output: [-1,-1,-1,-1]
Explanation: After the first day full lakes are [1]
After the second day full lakes are [1,2]
After the third day full lakes are [1,2,3]
After the fourth day full lakes are [1,2,3,4]
There's no day to dry any lake and there is no flood in any lake.
```

Example 2:

```
Input: rains = [1,2,0,0,2,1]
Output: [-1,-1,2,1,-1,-1]
Explanation: After the first day full lakes are [1]
After the second day full lakes are [1,2]
After the third day, we dry lake 2. Full lakes are [1]
After the fourth day, we dry lake 1. There is no full lakes.
After the fifth day, full lakes are [2].
After the sixth day, full lakes are [1,2].
It is easy that this scenario is flood-free. [-1,-1,1,2,-1,-1] is another acceptable scenario.
```

Example 3:

```
Input: rains = [1,2,0,1,2]
Output: []
Explanation: After the second day, full lakes are  [1,2]. We have to dry one lake in the third day.
After that, it will rain over lakes [1,2]. It's easy to prove that no matter which lake you choose to dry in the 3rd day, the other one will flood.
```

Example 4:

```
Input: rains = [69,0,0,0,69]
Output: [-1,69,1,1,-1]
Explanation: Any solution on one of the forms [-1,69,x,y,-1], [-1,x,69,y,-1] or [-1,x,y,69,-1] is acceptable where 1 <= x,y <= 10^9
```

Example 5:

```
Input: rains = [10,20,20]
Output: []
Explanation: It will rain over lake 20 two consecutive days. There is no chance to dry any lake.
```

Constraints:

1 <= rains.length <= 10^5  
0 <= rains[i] <= 10^9

## 方法

```go
func avoidFlood(rains []int) []int {
    var zeros []int
    res := make([]int, len(rains))
    m := make(map[int]int)
    
    for i, r := range rains {        
        if r == 0 {
            zeros = append(zeros, i)
            continue
        }
        
        if _, ok := m[r]; ok {
            zi := -1
            
			// find the closest index of zeros that could dry the current fulled lake
			//  O(NlogN) if using binary search to find the closest index
            for ti, tzi := range zeros {
                if tzi > m[r] {
                    zi = tzi
                    zeros = append(zeros[:ti], zeros[ti+1:] ...)
                    break
                }
            }
            
            if zi == -1 {
                return []int{}
            }
            
            res[zi] = r
        } 
        
        m[r] = i
        res[i] = -1
    }
    
	// fill the unused zeros index
    for i := range res {
        if res[i] == 0 {
            res[i] = 1
        }
    }
    
    return res
}
```


```python
class Solution:
    def avoidFlood(self, rains: List[int]) -> List[int]:
        ans = [1] * len(rains)
        left = []
        record = dict()
        for i, val in enumerate(rains):
            if val > 0:
                ans[i] = -1
                if val in record:
                    pos = bisect_left(left, record[val])
                    if pos >= len(left):
                        return []
                    ans[left.pop(pos)] = val
                record[val] = i
            else:
                left.append(i)
        return ans
```