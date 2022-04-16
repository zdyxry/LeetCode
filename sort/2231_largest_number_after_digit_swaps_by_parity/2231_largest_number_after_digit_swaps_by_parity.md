2231. Largest Number After Digit Swaps by Parity


Easy


You are given a positive integer num. You may swap any two digits of num that have the same parity (i.e. both odd digits or both even digits).

Return the largest possible value of num after any number of swaps.

 

Example 1:

```
Input: num = 1234
Output: 3412
Explanation: Swap the digit 3 with the digit 1, this results in the number 3214.
Swap the digit 2 with the digit 4, this results in the number 3412.
Note that there may be other sequences of swaps but it can be shown that 3412 is the largest possible number.
Also note that we may not swap the digit 4 with the digit 1 since they are of different parities.
```

Example 2:

```
Input: num = 65875
Output: 87655
Explanation: Swap the digit 8 with the digit 6, this results in the number 85675.
Swap the first digit 5 with the digit 7, this results in the number 87655.
Note that there may be other sequences of swaps but it can be shown that 87655 is the largest possible number.
```

Constraints:

1 <= num <= 109


## 方法



```python
class Solution:
    def largestInteger(self, num: int) -> int:
        even = []
        odd = []
        for idx, v in enumerate(str(num)):
            if int(v) % 2 == 0:
                even.append(v)
            if int(v) % 2 == 1:
                odd.append(v)
        even.sort(reverse=True)
        odd.sort(reverse=True)
        result = []
        for i in str(num):
            if int(i) % 2 == 0:
                result.append(even[0])
                even = even[1:]
            else:
                result.append(odd[0])
                odd = odd[1:]
        return int(''.join(result))
```


```go

func largestInteger(num int) int {
	s := []byte(strconv.Itoa(num))
	var a []byte
	var b []byte
	for i := range s {
		if (s[i]-'0')&1 == 1 {
			a = append(a, s[i])
		} else {
			b = append(b, s[i])
		}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})
	for i := range s {
		if (s[i]-'0')&1 == 1 {
			s[i] = a[0]
			a = a[1:]
		} else {
			s[i] = b[0]
			b = b[1:]
		}
	}
	r, _ := strconv.Atoi(string(s))
	return r
}

```