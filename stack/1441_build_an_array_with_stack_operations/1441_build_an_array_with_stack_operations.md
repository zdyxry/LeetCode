1441. Build an Array With Stack Operations


Easy


Given an array target and an integer n. In each iteration, you will read a number from  list = {1,2,3..., n}.

Build the target array using the following operations:

```
Push: Read a new element from the beginning list, and push it in the array.
Pop: delete the last element of the array.
If the target array is already built, stop reading more elements.
You are guaranteed that the target array is strictly increasing, only containing numbers between 1 to n inclusive.
```

Return the operations to build the target array.

You are guaranteed that the answer is unique.

 

Example 1:

```
Input: target = [1,3], n = 3
Output: ["Push","Push","Pop","Push"]
Explanation: 
Read number 1 and automatically push in the array -> [1]
Read number 2 and automatically push in the array then Pop it -> [1]
Read number 3 and automatically push in the array -> [1,3]
```

Example 2:

```
Input: target = [1,2,3], n = 3
Output: ["Push","Push","Push"]
```

Example 3:

```
Input: target = [1,2], n = 4
Output: ["Push","Push"]
Explanation: You only need to read the first 2 numbers and stop.
```

Example 4:

```
Input: target = [2,3,4], n = 4
Output: ["Push","Pop","Push","Push","Push"]
```
 

Constraints:

1 <= target.length <= 100  
1 <= target[i] <= 100  
1 <= n <= 100  
target is strictly increasing.


## 方法

```go
func buildArray(target []int, n int) []string {
    var (
		ret         []string
		pushedIndex int = 0
	)

	for i := 0; i < len(target); i++ {

		if target[i] == pushedIndex+1 {
			ret = append(ret, `Push`)
			pushedIndex++

		}

		if target[i] > pushedIndex+1 {

			for j := 0; j < target[i]-pushedIndex-1; j++ {
				ret = append(ret, `Push`)
				ret = append(ret, `Pop`)
			}

			pushedIndex += target[i] - pushedIndex - 1

			ret = append(ret, `Push`)

			pushedIndex++

		}

	}

	return ret

}
```


```python
class Solution:
    def buildArray(self, target: List[int], n: int) -> List[str]:
        res = []
        keep = 0
        for num in range(1, n+1):
            res.append("Push")
            keep += 1
            if num not in target:
                res.append("Pop")
                keep -= 1
            if keep == len(target):
                break
        return res
```