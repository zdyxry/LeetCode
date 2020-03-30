1395. Count Number of Teams


Medium


There are n soldiers standing in a line. Each soldier is assigned a unique rating value.

You have to form a team of 3 soldiers amongst them under the following rules:

Choose 3 soldiers with index (i, j, k) with rating (rating[i], rating[j], rating[k]).

A team is valid if:  (rating[i] < rating[j] < rating[k]) or (rating[i] > rating[j] > rating[k]) where (0 <= i < j < k < n).

Return the number of teams you can form given the conditions. (soldiers can be part of multiple teams).

 

Example 1:

```
Input: rating = [2,5,3,4,1]
Output: 3
Explanation: We can form three teams given the conditions. (2,3,4), (5,4,1), (5,3,1). 
```

Example 2:

```
Input: rating = [2,1,3]
Output: 0
Explanation: We can't form any team given the conditions.
```

Example 3:

```
Input: rating = [1,2,3,4]
Output: 4
```
 

Constraints:

n == rating.length  
1 <= n <= 200  
1 <= rating[i] <= 10^5


## 方法




```go
func numTeams(rating []int) int {
    rLen := len(rating)
    count := 0
    for i := 1; i < rLen - 1; i++ {
        count1, count2 := 0, 0
        for j := 0; j < i; j++ {
            if rating[j] < rating[i] {
                count1++
            }
        }
        for j := i + 1; j < rLen; j++ {
            if rating[j] < rating[i] {
                count2++
            }
        }
        count += count1 * (rLen - 1 - i - count2) + (i - count1) * count2;
    }
    return count
}
```


```python
class Solution(object):
    def numTeams(self, rating):
        """
        :type rating: List[int]
        :rtype: int
        """
        n = len(rating)
        ans = 0
        # 枚举每个士兵作为中间
        for i in range(1,n-1):
            #print(rating[i])
            # 左小右大
            l1,r1= 0,0
            # 左大右小
            l2,r2 =0,0
            for j in range(i-1,-1,-1):
                if rating[j] < rating[i]:
                    l1 += 1
                else:
                    l2 += 1
            for j in range(i+1,n):
                if rating[j] > rating[i]:
                    r1 += 1
                else:
                    r2 += 1
            #print(ans)
            # 排列组合
            ans += l1*r1 + l2*r2
        return ans

```