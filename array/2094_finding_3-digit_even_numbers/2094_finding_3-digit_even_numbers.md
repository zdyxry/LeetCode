2094. Finding 3-Digit Even Numbers


Easy


You are given an integer array digits, where each element is a digit. The array may contain duplicates.

You need to find all the unique integers that follow the given requirements:

```
The integer consists of the concatenation of three elements from digits in any arbitrary order.
The integer does not have leading zeros.
The integer is even.
For example, if the given digits were [1, 2, 3], integers 132 and 312 follow the requirements.
```

Return a sorted array of the unique integers.

 

Example 1:

```
Input: digits = [2,1,3,0]
Output: [102,120,130,132,210,230,302,310,312,320]
Explanation: All the possible integers that follow the requirements are in the output array. 
Notice that there are no odd integers or integers with leading zeros.
```

Example 2:

```
Input: digits = [2,2,8,8,2]
Output: [222,228,282,288,822,828,882]
Explanation: The same digit can be used as many times as it appears in digits. 
In this example, the digit 8 is used twice each time in 288, 828, and 882. 
```

Example 3:

```
Input: digits = [3,7,5]
Output: []
Explanation: No even integers can be formed using the given digits.
```

Constraints:

```
3 <= digits.length <= 100
0 <= digits[i] <= 9
```


## 方法


```
func findEvenNumbers(digits []int) []int {
    f := make([]int, 10)
    for _, n := range digits {
        f[n]++
    }
    
    r := make([]int, 0)
    for a := 1; a <= 9; a++ {
        if f[a] == 0 {
            continue
        }
        f[a]--
        
        for b := 0; b <= 9; b++ {
            if f[b] == 0 {
                continue
            }
            
            f[b]--
            for c := 0; c <= 8; c += 2 {
                if f[c] > 0 {
                    r = append(r, 100*a + 10*b + c)
                }
            }
            f[b]++
        }
        f[a]++
    }
    return r
}
```


```
class Solution:
    def findEvenNumbers(self, digits: List[int]) -> List[int]:
        all_elements = [i for i in range(100, 1000) if i%2==0]
        result = []
        
        for i in all_elements:
            current_i = [int(x) for x in str(i)]
            digits_copy = digits.copy()
            
            for x in range(3):
                if current_i[x] not in digits_copy:
                    break
                else:
                    digits_copy.remove(current_i[x])
                    
                if x == 2:
                    result.append(i) 
        
        
        return sorted(result)

```