1130. Minimum Cost Tree From Leaf Values


Medium


Given an array arr of positive integers, consider all binary trees such that:

```
Each node has either 0 or 2 children;
The values of arr correspond to the values of each leaf in an in-order traversal of the tree.  (Recall that a node is a leaf if and only if it has 0 children.)
The value of each non-leaf node is equal to the product of the largest leaf value in its left and right subtree respectively.
Among all possible binary trees considered, return the smallest possible sum of the values of each non-leaf node.  It is guaranteed this sum fits into a 32-bit integer.
```
 

Example 1:

```
Input: arr = [6,2,4]
Output: 32
Explanation:
There are two possible trees.  The first has non-leaf node sum 36, and the second has non-leaf node sum 32.

    24            24
   /  \          /  \
  12   4        6    8
 /  \               / \
6    2             2   4
```

Constraints:

2 <= arr.length <= 40   
1 <= arr[i] <= 15   
It is guaranteed that the answer fits into a 32-bit signed integer (ie. it is less than 2^31).


## 方法

```go
func mctFromLeafValues(arr []int) int {
    stack := []int{math.MaxInt32}
    
    res := 0
    
    for _, v := range arr{
        for stack[len(stack)-1] <= v{
            num := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            res = res + num * int(math.Min(float64(v), float64(stack[len(stack)-1])))
        }
        stack = append(stack, v)
    }
    
    for len(stack) > 2{
        num := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = res + num * stack[len(stack)-1]
    }
    return res
}
```


```python
class Solution:
    def mctFromLeafValues(self, arr: List[int]) -> int:
        res, n = 0, len(arr)
        stack = [float('inf')]
        for a in arr:
            while stack[-1] <= a:
                mid = stack.pop()
                res += mid * min(stack[-1], a)
            stack.append(a)
        while len(stack) > 2:
            res += stack.pop() * stack[-1]
        return res
```