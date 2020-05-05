1432. Max Difference You Can Get From Changing an Integer


Medium


You are given an integer num. You will apply the following steps exactly two times:

Pick a digit x (0 <= x <= 9).  
Pick another digit y (0 <= y <= 9). The digit y can be equal to x.  
Replace all the occurrences of x in the decimal representation of num by y.  
The new integer cannot have any leading zeros, also the new integer cannot be 0.  
Let a and b be the results of applying the operations to num the first and second times, respectively.

Return the max difference between a and b.

 

Example 1:

```
Input: num = 555
Output: 888
Explanation: The first time pick x = 5 and y = 9 and store the new integer in a.
The second time pick x = 5 and y = 1 and store the new integer in b.
We have now a = 999 and b = 111 and max difference = 888
```

Example 2:

```
Input: num = 9
Output: 8
Explanation: The first time pick x = 9 and y = 9 and store the new integer in a.
The second time pick x = 9 and y = 1 and store the new integer in b.
We have now a = 9 and b = 1 and max difference = 8
```

Example 3:

```
Input: num = 123456
Output: 820000
```

Example 4:

```
Input: num = 10000
Output: 80000
```

Example 5:
```
Input: num = 9288
Output: 8700
```

Constraints:

1 <= num <= 10^8

## 方法

```go
func toDigits(num int) []int {
    if num < 10 {
        return []int{num}
    }
    return append(toDigits(num/10), num %10)
}


func maxDiff(num int) int {
    digits := toDigits(num)
    
    digitA, digitB := -1, -1
    var zero bool
    for _, d := range digits {
        if digitA < 0 && d < 9 {
            digitA = d
        }
        if digitB < 0 && d >= 1 {
            if d == digits[0] {
                if d > 1 {
                    digitB = d
                }
            } else {
                digitB = d
                zero = true
            }
        }
        if digitA >= 0 && digitB >= 0 {
            break
        }
    }
    a, b := 0, 0
    for _, d := range digits {
        a *= 10
        b *= 10
        if d == digitA {
            a += 9
        } else {
            a += d
        }
        if d == digitB {
            if !zero {
                b++
            }
        } else {
            b += d
        }
    }
    return a - b
    
}
```


```python
class Solution:
    def maxDiff(self, num: int) -> int:
        a = b = str(num)

        for digit in a:
            if digit != "9":
                a = a.replace(digit, "9")
                break
        
        if b[0] != "1":
            b = b.replace(b[0], "1")
        else:
            for digit in b[1:]:
                if digit not in "01":
                    b = b.replace(digit, "0")
                    break
        
        return int(a) - int(b)
```
