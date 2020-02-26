1363. Largest Multiple of Three


Hard


Given an integer array of digits, return the largest multiple of three that can be formed by concatenating some of the given digits in any order.

Since the answer may not fit in an integer data type, return the answer as a string.

If there is no answer return an empty string.

 

Example 1:

```
Input: digits = [8,1,9]
Output: "981"
```

Example 2:

```
Input: digits = [8,6,7,1,0]
Output: "8760"
```

Example 3:

```
Input: digits = [1]
Output: ""
```

Example 4:

```
Input: digits = [0,0,0,0,0,0]
Output: "0"
```
 

Constraints:

1 <= digits.length <= 10^4  
0 <= digits[i] <= 9  
The returning answer must not contain unnecessary leading zeros.


## 方法

```go
import (
    "fmt"
    "sort"
    "strings"
    "strconv"
)


type sortInt []int 
func (a sortInt) Len() int           { return len(a) }
func (a sortInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortInt) Less(i, j int) bool { return a[i] > a[j] } //reversed

func largestMultipleOfThree(digits []int) string {

    arr := sortInt(digits)
    sort.Sort(arr)
    
    var li0 []*int
    var li1 []*int
    var li2 []*int
    
    total := 0
    
    cnt1 := 0
    cnt2 := 0
    for i, v := range arr {
        switch v%3 {
            case 0:
                li0 = append(li0, &arr[i])
                total += v
            case 1:
                li1 = append(li1, &arr[i])
                cnt1 += 1
                total += v
            case 2:
                li2 = append(li2, &arr[i])
                cnt2 += 1
                total += v
        }
    }
    if total % 3 == 0 {
        //use all
    } else if total % 3 == 1 {
        if len(li1) >= 1 {
            //remove 1 li1
            p := li1[len(li1)-1]
            *p = -10
        } else if len(li2) >= 2 {
            //remove 2 li2
            p := li2[len(li2)-1]
            *p = -10
            p = li2[len(li2)-2]
            *p = -10
        } else {
            //remove 1 li2 if exists
            if len(li2) >= 1 {
                p := li2[len(li2)-1]
                *p = -10
            }
        }
    } else if total %3 == 2 {
        if len(li2) >= 1 {
            //remove 1 li2
            p := li2[len(li2)-1]
            *p = -10
        } else if len(li1) >= 2 {
            //remove 2 li1
            p := li1[len(li1)-1]
            *p = -10
            p = li1[len(li1)-2]
            *p = -10
        } else {
            //remove 1 li1 if exists
            if len(li1) >= 1 {
                p := li1[len(li1)-1]
                *p = -10
            }
        }
    }
    //fmt.Println(arr)
    var out []string
    for _, num := range arr {
        if num >= 0 {
            s := strconv.Itoa(num)
            out = append(out, s)
        }
    }
        
    tmp := strings.Join(out, "")
    if tmp == "" {return ""}
    n, _ := strconv.Atoi(tmp)
    if n == 0 { return "0"}

    return tmp
}
```




```python
class Solution(object):
    def largestMultipleOfThree(self, A):
        """
        :type digits: List[int]
        :rtype: str
        """
        total = sum(A)
        count = collections.Counter(A)
        A.sort(reverse=1)

        def f(i):
            if count[i]:
                A.remove(i)
                count[i] -= 1
            if not A: return ''
            if not any(A): return '0'
            if sum(A) % 3 == 0: return ''.join(map(str, A))

        if total % 3 == 0:
            return f(-1)
        if total % 3 == 1 and count[1] + count[4] + count[7]:
            return f(1) or f(4) or f(7)
        if total % 3 == 2 and count[2] + count[5] + count[8]:
            return f(2) or f(5) or f(8)
        if total % 3 == 2:
            return f(1) or f(1) or f(4) or f(4) or f(7) or f(7)
        return f(2) or f(2) or f(5) or f(5) or f(8) or f(8)
```