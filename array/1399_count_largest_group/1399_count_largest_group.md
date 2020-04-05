1399. Count Largest Group


Easy


Given an integer n. Each number from 1 to n is grouped according to the sum of its digits. 

Return how many groups have the largest size.

 

Example 1:

```
Input: n = 13
Output: 4
Explanation: There are 9 groups in total, they are grouped according sum of its digits of numbers from 1 to 13:
[1,10], [2,11], [3,12], [4,13], [5], [6], [7], [8], [9]. There are 4 groups with largest size.
```

Example 2:

```
Input: n = 2
Output: 2
Explanation: There are 2 groups [1], [2] of size 1.
```

Example 3:

```
Input: n = 15
Output: 6
```

Example 4:

```
Input: n = 24
Output: 5
```
 

Constraints:

1 <= n <= 10^4


## 方法


```go
func countLargestGroup(n int) int {
    var (
        i     int
        total []int
        countMap = map[int]int{}
    )

    for i = 1; i <= n; i++ {
        countMap[count(i)] ++
    }
    for _, v := range countMap {
        total= append(total,v)
    }
    sort.Slice(total, func(i, j int) bool {
        return total[i] > total[j]
    })

    for i = 1; i< len(total); i++ {
        if total[i] != total[0] {
            return i
        }
    }

    return len(total)
}

func count(n int) (rst int) {
    for n > 0 {
        rst += (n % 10)
        n /= 10
    }
    return
}
```


```python

class Solution:
    def countLargestGroup(self, n: int) -> int:
        d = {}
        def dsum(num):
            s = 0
            while num:
                s += num % 10
                num = num // 10
            return s
        
        for i in range(1, n+1):
            ds = dsum(i)
            if ds not in d:
                d[ds] = 1
            else:
                d[ds] += 1
        return len([v for v in d.values() if v == max(d.values())])
```