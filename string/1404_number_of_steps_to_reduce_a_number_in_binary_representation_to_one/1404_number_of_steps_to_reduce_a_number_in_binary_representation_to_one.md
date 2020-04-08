1404. Number of Steps to Reduce a Number in Binary Representation to One


Medium


Given a number s in their binary representation. Return the number of steps to reduce it to 1 under the following rules:

If the current number is even, you have to divide it by 2.

If the current number is odd, you have to add 1 to it.

It's guaranteed that you can always reach to one for all testcases.

 

Example 1:

```
Input: s = "1101"
Output: 6
Explanation: "1101" corressponds to number 13 in their decimal representation.
Step 1) 13 is odd, add 1 and obtain 14. 
Step 2) 14 is even, divide by 2 and obtain 7.
Step 3) 7 is odd, add 1 and obtain 8.
Step 4) 8 is even, divide by 2 and obtain 4.  
Step 5) 4 is even, divide by 2 and obtain 2. 
Step 6) 2 is even, divide by 2 and obtain 1.  
```

Example 2:

```
Input: s = "10"
Output: 1
Explanation: "10" corressponds to number 2 in their decimal representation.
Step 1) 2 is even, divide by 2 and obtain 1.  
```

Example 3:

```
Input: s = "1"
Output: 0
```
 

Constraints:

1 <= s.length <= 500  
s consists of characters '0' or '1'  
s[0] == '1'


## 方法


```go
func reverse(A []byte) {
	for i := 0; i < len(A)/2; i++ {
		A[i], A[len(A)-1-i] = A[len(A)-1-i], A[i]
	}
}

func numSteps(s string) int {
	if len(s) == 1 {
		return 0
	}
	bs := []byte(s)
	reverse(bs)
	step, carry, last := 0, 0, bs[len(bs)-1]
	bs = bs[:len(bs)-1]
	for _, v := range bs {
		if carry == 0 {
			if v == '1' {
				step += 2
				carry = 1
			} else {
				step += 1
			}
		} else {
			if v == '0' {
				step += 2
			} else {
				step += 1
			}
		}
	}
	if last == '1' {
		step += carry
	}
	return step
}
```


```python
class Solution:
    def numSteps(self, s: str) -> int:
        cnt = 0
        s_num = int(s, 2)
        while s_num != 1:
            if s_num % 2 == 1:
                s_num += 1
            else:
                s_num = s_num // 2
            cnt += 1
        return cnt
```