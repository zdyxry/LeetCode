1089. Duplicate Zeros


Easy


Given a fixed length array arr of integers, duplicate each occurrence of zero, shifting the remaining elements to the right.

Note that elements beyond the length of the original array are not written.

Do the above modifications to the input array in place, do not return anything from your function.

 

Example 1:

```
Input: [1,0,2,3,0,4,5,0]
Output: null
Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]
```

Example 2:

```
Input: [1,2,3]
Output: null
Explanation: After calling your function, the input array is modified to: [1,2,3]
```

Note:

1 <= arr.length <= 10000  
0 <= arr[i] <= 9

## 方法


```go
func duplicateZeros(arr []int)  {
    for index := 0; index < len(arr); index++ {
		if arr[index] == 0 && index+1 < len(arr) {
			arr = append(arr[:index+1], arr[index:len(arr)-1]...)
			index++
		}
	}
}
```


```python
class Solution:
    def duplicateZeros(self, arr: List[int]) -> None:
        """
        Do not return anything, modify arr in-place instead.
        """
        nums = []
        for num in arr:
            nums.append(num)
            if num == 0:
                nums.append(0)
        for i in range(len(arr)):
            arr[i] = nums[i]
```