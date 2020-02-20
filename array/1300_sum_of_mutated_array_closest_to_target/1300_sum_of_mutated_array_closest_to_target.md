1300. Sum of Mutated Array Closest to Target


Medium


Given an integer array arr and a target value target, return the integer value such that when we change all the integers larger than value in the given array to be equal to value, the sum of the array gets as close as possible (in absolute difference) to target.

In case of a tie, return the minimum such integer.

Notice that the answer is not neccesarilly a number from arr.

 

Example 1:

```
Input: arr = [4,9,3], target = 10
Output: 3
Explanation: When using 3 arr converts to [3, 3, 3] which sums 9 and that's the optimal answer.
```

Example 2:

```
Input: arr = [2,3,5], target = 10
Output: 5
```

Example 3:

```
Input: arr = [60864,25176,27249,21296,20204], target = 56803
Output: 11361
```

 

Constraints:

1 <= arr.length <= 10^4

1 <= arr[i], target <= 10^5


## 方法

```go
func findBestValue(arr []int, target int) int {
    right := 0
	for i := range arr {
		if arr[i] > right {
			right = arr[i]
		}
	}
	right += 1
	dif := math.MaxUint16
	res := arr[0]
	left := 0
	for left < right {
		mid := (left + right) / 2
		sum := 0
		for _, val := range arr {
			if val > mid {
				sum += mid
			} else {
				sum += val
			}
		}
		temp := int(math.Abs(float64(sum - target)))
		if temp < dif {
			dif = temp
			res = mid
		} else if temp == dif {
			if res > mid {
				res = mid
			}
		}
		if sum < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return res
}
```




```python
class Solution(object):
    def findBestValue(self, arr, target):
        """
        :type arr: List[int]
        :type target: int
        :rtype: int
        """
        arr.sort(reverse = True)
        
        while arr and target >= arr[-1]*len(arr):
            temp = arr[-1]
            target -= arr.pop()
        
        if not arr:
            return temp
        res = target/float(len(arr))
        
        if res % 1 > 0.5:
            return int(res)+1
        else:
            return int(res)
```