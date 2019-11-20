842. Split Array into Fibonacci Sequence


Medium


Given a string S of digits, such as S = "123456579", we can split it into a Fibonacci-like sequence [123, 456, 579].

Formally, a Fibonacci-like sequence is a list F of non-negative integers such that:

0 <= F[i] <= 2^31 - 1, (that is, each integer fits a 32-bit signed integer type);
F.length >= 3;
and F[i] + F[i+1] = F[i+2] for all 0 <= i < F.length - 2.
Also, note that when splitting the string into pieces, each piece must not have extra leading zeroes, except if the piece is the number 0 itself.

Return any Fibonacci-like sequence split from S, or return [] if it cannot be done.

Example 1:

```
Input: "123456579"
Output: [123,456,579]
```

Example 2:

```
Input: "11235813"
Output: [1,1,2,3,5,8,13]
```

Example 3:

```
Input: "112358130"
Output: []
Explanation: The task is impossible.
```

Example 4:

```
Input: "0123"
Output: []
Explanation: Leading zeroes are not allowed, so "01", "2", "3" is not valid.
```

Example 5:

```
Input: "1101111"
Output: [110, 1, 111]
Explanation: The output [11, 0, 11, 11] would also be accepted.
```

Note:

1 <= S.length <= 200  
S contains only digits.

## 方法

```go
func splitIntoFibonacci(S string) []int {
    if len(S) < 3 {
		return []int{}
	}
	res, isComplete := []int{}, false
	for firstEnd := 0; firstEnd < len(S)/2; firstEnd++ {
		if S[0] == '0' && firstEnd > 0 {
			break
		}
		first, _ := strconv.Atoi(S[:firstEnd+1])
		if first >= 1<<31 { // 题目要求每个数都要小于 2^31 - 1 = 2147483647，此处剪枝很关键！
			break
		}
		for secondEnd := firstEnd + 1; max(firstEnd, secondEnd-firstEnd) <= len(S)-secondEnd; secondEnd++ {
			if S[firstEnd+1] == '0' && secondEnd-firstEnd > 1 {
				break
			}
			second, _ := strconv.Atoi(S[firstEnd+1 : secondEnd+1])
			if second >= 1<<31 { // 题目要求每个数都要小于 2^31 - 1 = 2147483647，此处剪枝很关键！
				break
			}
			findRecursiveCheck(S, first, second, secondEnd+1, &res, &isComplete)
		}
	}
	return res
}

//Propagate for rest of the string
func findRecursiveCheck(S string, x1 int, x2 int, left int, res *[]int, isComplete *bool) {
	if x1 >= 1<<31 || x2 >= 1<<31 { // 题目要求每个数都要小于 2^31 - 1 = 2147483647，此处剪枝很关键！
		return
	}
	if left == len(S) {
		if !*isComplete {
			*isComplete = true
			*res = append(*res, x1)
			*res = append(*res, x2)
		}
		return
	}
	if strings.HasPrefix(S[left:], strconv.Itoa(x1+x2)) && !*isComplete {
		*res = append(*res, x1)
		findRecursiveCheck(S, x2, x1+x2, left+len(strconv.Itoa(x1+x2)), res, isComplete)
		return
	}
	if len(*res) > 0 && !*isComplete {
		*res = (*res)[:len(*res)-1]
	}
	return
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
```


```python
class Solution(object):
    def splitIntoFibonacci(self, S):
        """
        :type S: str
        :rtype: List[int]
        """
        res=[]
        current=[]
        def backtrack(cursor,current):     #``cursor'' represents the current scanning cursor, ``current'' represents the current partial result
            if len(current)>=3 and cursor==len(S):        # reach a result
                res.append(current)
                return
            if len(current)<2:                           #add first two items into partial result
                for index in range(cursor,len(S)):
                    avai=S[cursor:index+1]
                    if (avai[0]=="0" and len(avai)>1) or int(avai)>2**31-1:    # if have leading zeros, break
                        return
                    backtrack(index+1,current+[int(avai)])
            else:
                a=current[-1]
                b=current[-2]
                seek=str(a+b)
                if S[cursor:cursor+len(seek)]==seek and int(seek)<2**31-1:
                    backtrack(cursor+len(seek),current+[int(seek)])
                else:
                    return
        backtrack(0,current)
        if res:
            return res[0]
        else:
            return res
```