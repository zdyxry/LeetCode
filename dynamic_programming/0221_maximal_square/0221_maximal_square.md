221. Maximal Square

Medium

Given a 2D binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.

Example:


```
Input: 

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

Output: 4
```

## 方法


```go
func maximalSquare(matrix [][]byte) int {
    if len(matrix)==0{
        return 0
    }
    dp := make([][]int, 2)
    dp[0]=make([]int, len(matrix[0]))
    dp[1]=make([]int, len(matrix[0]))
    maxLen := 0
    min := func(x, y int) int{
        if x<y{
            return x
        }
        return y
    }
    for i := range matrix[0]{
        dp[0][i] = int(matrix[0][i]-'0')
        dp[1][i] = int(matrix[0][i]-'0')
        maxLen = maxLen + dp[1][i] - min(maxLen,dp[1][i])
    }
    fmt.Println(maxLen)
    for i:= 1; i < len(matrix); i++ {
        dp[1][0]=int(matrix[i][0]-'0')
        maxLen = maxLen + dp[1][0] - min(maxLen,dp[1][0])
        for j := 1;j<len(matrix[0]);j++{
            dp[1][j]=0
            if int(matrix[i][j]-'0') == 1{
                dp[1][j]=1+min(dp[0][j-1],min(dp[0][j],dp[1][j-1]))
            }
            maxLen = maxLen + dp[1][j] - min(maxLen, dp[1][j])
        }
        for i,val:=range dp[1]{
            dp[0][i]=val
        }
    } 
    return maxLen*maxLen
}
```


```python
class Solution(object):
    def maximalSquare(self, matrix):
        """
        :type matrix: List[List[str]]
        :rtype: int
        """
        if not matrix or not matrix[0]:
            return 0
        rows,cols=len(matrix),len(matrix[0])
        max_side=0
        square_sides=[0]*cols
        for r in range(rows):
            new_square_sides=[int(matrix[r][0])] + [0 for _ in range(cols-1)]
            for c in range(1,cols):
                if matrix[r][c]=="1":
                    new_square_sides[c]=1+min(new_square_sides[c-1],square_sides[c],square_sides[c-1])
            max_side=max(max_side,max(new_square_sides))
            square_sides=new_square_sides
        return max_side**2 
```
