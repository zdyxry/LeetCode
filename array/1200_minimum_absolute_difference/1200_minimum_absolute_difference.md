1200. Minimum Absolute Difference


Easy


Given an array of distinct integers arr, find all pairs of elements with the minimum absolute difference of any two elements. 

Return a list of pairs in ascending order(with respect to pairs), each pair [a, b] follows

```
a, b are from arr
a < b
b - a equals to the minimum absolute difference of any two elements in arr
```

Example 1:

```
Input: arr = [4,2,1,3]
Output: [[1,2],[2,3],[3,4]]
Explanation: The minimum absolute difference is 1. List all pairs with difference equal to 1 in ascending order.
```

Example 2:

```
Input: arr = [1,3,6,10,15]
Output: [[1,3]]
```

Example 3:

```
Input: arr = [3,8,-10,23,19,-4,-14,27]
Output: [[-14,-10],[19,23],[23,27]]
```

Constraints:

2 <= arr.length <= 10^5  
-10^6 <= arr[i] <= 10^6

## 方法

```go
func minimumAbsDifference(arr []int) [][]int {
    min := math.MaxInt32
	res := [][]int{}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] < min {
			min = arr[i] - arr[i-1]
		}
		if min == 1 {
			break
		}
	}
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] == min {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}
	return res
}
```


```python
class Solution:
    def minimumAbsDifference(self, arr: List[int]) -> List[List[int]]:
        arr.sort()
        m = float('inf')
        out = []
        for i in range(1, len(arr)):
            prev = arr[i - 1]
            curr = abs(prev - arr[i])
            if curr < m:
                out = [[prev, arr[i]]]
                m = curr
            elif curr == m: out.append([prev, arr[i]])
        return out
```