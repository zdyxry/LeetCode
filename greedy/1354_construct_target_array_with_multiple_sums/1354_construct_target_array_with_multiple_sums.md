1354. Construct Target Array With Multiple Sums


Hard


Given an array of integers target. From a starting array, A consisting of all 1's, you may perform the following procedure :

let x be the sum of all elements currently in your array.

choose index i, such that 0 <= i < target.size and set the value of A at index i to x.

You may repeat this procedure as many times as needed.

Return True if it is possible to construct the target array from A otherwise return False.

 

Example 1:

```
Input: target = [9,3,5]
Output: true
Explanation: Start with [1, 1, 1] 
[1, 1, 1], sum = 3 choose index 1
[1, 3, 1], sum = 5 choose index 2
[1, 3, 5], sum = 9 choose index 0
[9, 3, 5] Done
```

Example 2:

```
Input: target = [1,1,1,2]
Output: false
Explanation: Impossible to create target array from [1,1,1,1].
```

Example 3:

```
Input: target = [8,5]
Output: true
```

 

Constraints:

N == target.length

1 <= target.length <= 5 * 10^4

1 <= target[i] <= 10^9


## 方法

```go
func isPossible(target []int) bool {
    if len(target) == 0 {
        return true
    }
    for j := 0; j < len(target); j++ {
        sum := 0
        maxIdx := 0
        maxNum := target[0]
        for i := 0; i < len(target); i++ {
            if target[i] > maxNum {
                maxIdx = i
                maxNum = target[i]
            }
            sum += target[i]
        }
        if maxNum == 1 {
            return true
        }
        if sum - maxNum >= maxNum || (sum - maxNum) == 0{
            return false
        }
        target[maxIdx] = maxNum % (sum - maxNum)
    }
    return true
}
```


```python
class Solution(object):
    def isPossible(self, target):
        """
        :type target: List[int]
        :rtype: bool
        """
        while True: 
            current_x = max(target)
            if current_x == 1:   
                return True # 最大的数等于 1 了而且数组中无小于 1 的数，那么说明整个数组都是 1，返回 True
            idx = target.index(current_x)
            s = sum(target)
            inc = s - current_x
            if inc >= current_x or inc == 0: 
                return False # inc 大于等于 x，这样将会出现小于 1 的值，不是合法情况，返回 False
            print current_x, inc
            target[idx] = current_x % inc
```