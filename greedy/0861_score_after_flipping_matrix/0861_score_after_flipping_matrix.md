861. Score After Flipping Matrix


Medium


We have a two dimensional matrix A where each value is 0 or 1.

A move consists of choosing any row or column, and toggling each value in that row or column: changing all 0s to 1s, and all 1s to 0s.

After making any number of moves, every row of this matrix is interpreted as a binary number, and the score of the matrix is the sum of these numbers.

Return the highest possible score.

 

Example 1:

```
Input: [[0,0,1,1],[1,0,1,0],[1,1,0,0]]
Output: 39
Explanation:
Toggled to [[1,1,1,1],[1,0,0,1],[1,1,1,1]].
0b1111 + 0b1001 + 0b1111 = 15 + 9 + 15 = 39
```

Note:

1 <= A.length <= 20   
1 <= A[0].length <= 20   
A[i][j] is 0 or 1.


## 方法


```go
func matrixScore(A [][]int) int {

	low := len(A)
	col := len(A[0])

	for i := 0 ; i < low ; i ++ {
		if A[i][0] != 1 {
			for j := 0 ;j < col ; j++{
				if 	A[i][j] == 0 {
					A[i][j] = 1
				}else {
					A[i][j] = 0
				}
			}
		}
	}
	for i := 0; i < col ; i ++ {
		res := 0
		for j := 0 ; j < low ; j++ {
			if A[j][i] != 0 {
				res++
			}
		}
		re := int(math.Ceil( float64(low*1.0) / 2) )
		if res < re  {
			for j := 0 ; j < low ; j++ {
				if A[j][i] != 0 {
					A[j][i]= 0
				}else{
					A[j][i]= 1
				}
			}
		}
	}
	result := 0
	for i := 0 ; i < low ; i ++{
		for j := 0 ; j < col ; j ++ {
			if A[i][j] != 0 {
				result += int(math.Pow(float64(2),float64(col - j-1)))
			}

		}
	}

	return result
}

```

```python
class Solution:
    def matrixScore(self, A: List[List[int]]) -> int:
        row = len(A)
        col = len(A[0])



        #先把所有的第一位都变成1
        for i in range(row):
            if A[i][0]==0:
                for j in range(len(A[i])):
                    A[i][j]=1-A[i][j]
        #后面的列，0多的就转，1多的就保留
        for i in range(1,col):
            num0 = 0
            num1 = 0
            #统计个数
            for k in range(row):
                if A[k][i]==0:
                    num0 +=1
                else:
                    num1 +=1
            if num1<num0:
                for j in range(row):
                    A[j][i]=1-A[j][i]

        res = 0
        for i in range(len(A)):
            temp = ""
            for j in range(len(A[0])):
                temp+=str(A[i][j])
            res+=int(temp,2)
        return res 
```