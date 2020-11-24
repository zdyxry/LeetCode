1090. Largest Values From Labels


Medium

We have a set of items: the i-th item has value values[i] and label labels[i].

Then, we choose a subset S of these items, such that:

```
|S| <= num_wanted
For every label L, the number of items in S with label L is <= use_limit.
Return the largest possible sum of the subset S.
```
 

Example 1:

```
Input: values = [5,4,3,2,1], labels = [1,1,2,2,3], num_wanted = 3, use_limit = 1
Output: 9
Explanation: The subset chosen is the first, third, and fifth item.
```

Example 2:

```
Input: values = [5,4,3,2,1], labels = [1,3,3,3,2], num_wanted = 3, use_limit = 2
Output: 12
Explanation: The subset chosen is the first, second, and third item.
```

Example 3:

```
Input: values = [9,8,8,7,6], labels = [0,0,0,1,1], num_wanted = 3, use_limit = 1
Output: 16
Explanation: The subset chosen is the first and fourth item.
```

Example 4:

```
Input: values = [9,8,8,7,6], labels = [0,0,0,1,1], num_wanted = 3, use_limit = 2
Output: 24
Explanation: The subset chosen is the first, second, and fourth item.
```

Note:

1 <= values.length == labels.length <= 20000   
0 <= values[i], labels[i] <= 20000   
1 <= num_wanted, use_limit <= values.length


## 方法


```go
import "sort"

type item struct {
    v int
    l int
}

func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
    if len(values) == 0 || num_wanted == 0 {
        return 0
    }
    
    items := make([]item, len(values))
    for i := 0; i < len(values); i ++ {
        items[i] = item{values[i], labels[i]}
    }
    
    sort.Slice(items, func(i, j int)bool{
        return items[i].v > items[j].v
    })
    
    res := 0
    cnt := make(map[int]int)
    total := 0
    for i := 0; i < len(items); i++ {
        
        if cnt[items[i].l] == use_limit {
            continue
        }
        
        res += items[i].v
        cnt[items[i].l]++
        total++
        
        if total == num_wanted {
            break
        }
    }
    
    return res
}
```


```python
class Solution:
    def largestValsFromLabels(self, values: List[int], labels: List[int], num_wanted: int, use_limit: int) -> int:
        l = [(val, label) for val, label in zip(values, labels)]
        l.sort(reverse = True)

        c = Counter()
        cnt = 0
        sum = 0
        for val, label in l:
            if cnt == num_wanted:
                break

            if c[label] < use_limit:
                c[label] += 1
                sum += val
                cnt += 1

        return sum
```