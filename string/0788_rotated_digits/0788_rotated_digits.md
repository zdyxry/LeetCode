788. Rotated Digits


Easy


X is a good number if after rotating each digit individually by 180 degrees, we get a valid number that is different from X.  Each digit must be rotated - we cannot choose to leave it alone.

```
A number is valid if each digit remains a digit after rotation. 0, 1, and 8 rotate to themselves; 2 and 5 rotate to each other (on this case they are rotated in a different direction, in other words 2 or 5 gets mirrored); 6 and 9 rotate to each other, and the rest of the numbers do not rotate to any other number and become invalid.
```

Now given a positive number N, how many numbers X from 1 to N are good?

Example:

```
Input: 10
Output: 4
Explanation: 
There are four good numbers in the range [1, 10] : 2, 5, 6, 9.
Note that 1 and 10 are not good numbers, since they remain unchanged after rotating.
```

Note:

N  will be in range [1, 10000].


## 方法


```python
class Solution:
    def rotatedDigits(self, N: int) -> int:
        ans = 0
        # For each x in [1, N], check whether it's good
        for x in range(1, N+1):
            S = str(x)
            # Each x has only rotateable digits, and one of them
            # rotates to a different digit
            ans += (all(d not in '347' for d in S)
                    and any(d in '2569' for d in S))
        return ans
```


```go
func good(n int ,flag bool) bool {
	if n == 0 {
		return  flag
	}
	d := n % 10
	switch d {
	case 3,4,7:
		return  false
	case 0,1,8:
		return good(n/10,flag)
	}
    //2,5,6,9
	return  good(n/10,true)
}

func rotatedDigits(N int) int {
	flag := false
	ans := 0
	for i := 1 ; i <= N; i++{
		if good(i,flag){
			ans++
		}
	}
	return ans
}
```