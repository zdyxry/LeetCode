1442. Count Triplets That Can Form Two Arrays of Equal XOR


Medium


Given an array of integers arr.

We want to select three indices i, j and k where (0 <= i < j <= k < arr.length).

Let's define a and b as follows:

```
a = arr[i] ^ arr[i + 1] ^ ... ^ arr[j - 1]
b = arr[j] ^ arr[j + 1] ^ ... ^ arr[k]
Note that ^ denotes the bitwise-xor operation.
```

Return the number of triplets (i, j and k) Where a == b.

 

Example 1:

```
Input: arr = [2,3,1,6,7]
Output: 4
Explanation: The triplets are (0,1,2), (0,2,2), (2,3,4) and (2,4,4)
```

Example 2:

```
Input: arr = [1,1,1,1,1]
Output: 10
```

Example 3:

```
Input: arr = [2,3]
Output: 0
```

Example 4:

```
Input: arr = [1,3,5,7,9]
Output: 3
```

Example 5:

```
Input: arr = [7,11,12,9,5,2,7,17,22]
Output: 8
```
 

Constraints:

1 <= arr.length <= 300  
1 <= arr[i] <= 10^8


## 方法

```go
func countTriplets(arr []int) int {
    n := len(arr) 
    prefix := make([]int, n+1, n+1)
    prefix[0] = 0
    for i := 1; i <= n ; i++ {
        prefix[i] = arr[i-1]^prefix[i-1];
    }

    total := 0
    for i := 0; i < n; i++ {
        for j := i+1; j <= n; j++ {
            if prefix[i] == prefix[j] {
                total += j - i - 1
            }
        }
    }
    return total
}
```



```python
class Solution:
    def countTriplets(self, arr: List[int]) -> int:
        if len(arr) < 2:
            return 0

        count = 0

        for i in range(len(arr)-1):
            temp = arr[i]
            for j in range(i+1, len(arr)):
                temp = temp^arr[j]
                if temp == 0:
                    count += j-i

        return count
```