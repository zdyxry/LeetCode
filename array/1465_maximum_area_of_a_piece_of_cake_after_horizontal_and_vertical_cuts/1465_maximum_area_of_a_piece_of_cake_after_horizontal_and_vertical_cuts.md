1465. Maximum Area of a Piece of Cake After Horizontal and Vertical Cuts


Medium


Given a rectangular cake with height h and width w, and two arrays of integers horizontalCuts and verticalCuts where horizontalCuts[i] is the distance from the top of the rectangular cake to the ith horizontal cut and similarly, verticalCuts[j] is the distance from the left of the rectangular cake to the jth vertical cut.

Return the maximum area of a piece of cake after you cut at each horizontal and vertical position provided in the arrays horizontalCuts and verticalCuts. Since the answer can be a huge number, return this modulo 10^9 + 7.

 

Example 1:

![1](1465-1.png)

```
Input: h = 5, w = 4, horizontalCuts = [1,2,4], verticalCuts = [1,3]
Output: 4 
Explanation: The figure above represents the given rectangular cake. Red lines are the horizontal and vertical cuts. After you cut the cake, the green piece of cake has the maximum area.
```

Example 2:

![2](1465-2.png)

```
Input: h = 5, w = 4, horizontalCuts = [3,1], verticalCuts = [1]
Output: 6
Explanation: The figure above represents the given rectangular cake. Red lines are the horizontal and vertical cuts. After you cut the cake, the green and yellow pieces of cake have the maximum area.
```

Example 3:

```
Input: h = 5, w = 4, horizontalCuts = [3], verticalCuts = [3]
Output: 9
```

Constraints:

2 <= h, w <= 10^9  
1 <= horizontalCuts.length < min(h, 10^5)  
1 <= verticalCuts.length < min(w, 10^5)  
1 <= horizontalCuts[i] < h  
1 <= verticalCuts[i] < w  
It is guaranteed that all elements in horizontalCuts are distinct.  
It is guaranteed that all elements in verticalCuts are distinct.


## 方法

```go
func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
    maxH, maxV := 0, 0
	pre := 0
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	for i := 0; i < len(horizontalCuts); i++ {
		maxH = max(maxH, horizontalCuts[i] - pre)
		pre = horizontalCuts[i]
	}
	maxH = max(maxH, h - pre)
	pre = 0
	for i := 0; i < len(verticalCuts); i++ {
		maxV = max(maxV, verticalCuts[i] - pre)
		pre = verticalCuts[i]
	}
	maxV = max(maxV, w - pre)
	return (maxH * maxV) % 1000000007
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
```


```python
class Solution:
    def maxArea(self, h: int, w: int, horizontalCuts: List[int], verticalCuts: List[int]) -> int:
        height, width = 0, 0
        mod = 10 ** 9 + 7
        horizontalCuts.sort()
        verticalCuts.sort()
        h_len, v_len = len(horizontalCuts), len(verticalCuts)

        for i in range(1, h_len):
            height = max(height, horizontalCuts[i] - horizontalCuts[i - 1])
        height = max(height, horizontalCuts[0], h - horizontalCuts[h_len - 1])

        for i in range(1, v_len):
            width = max(width, verticalCuts[i] - verticalCuts[i - 1])
        width = max(width, verticalCuts[0], w - verticalCuts[v_len - 1])

        return (height * width) % mod
```