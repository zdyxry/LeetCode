1405. Longest Happy String


Medium



A string is called happy if it does not have any of the strings 'aaa', 'bbb' or 'ccc' as a substring.

Given three integers a, b and c, return any string s, which satisfies following conditions:

s is happy and longest possible.  
s contains at most a occurrences of the letter 'a', at most b occurrences of the letter 'b' and at most c occurrences of the letter 'c'.  
s will only contain 'a', 'b' and 'c' letters.  
If there is no such string s return the empty string "".  

 

Example 1:

```
Input: a = 1, b = 1, c = 7
Output: "ccaccbcc"
Explanation: "ccbccacc" would also be a correct answer.
```

Example 2:

```
Input: a = 2, b = 2, c = 1
Output: "aabbc"
```

Example 3:

```
Input: a = 7, b = 1, c = 0
Output: "aabaa"
Explanation: It's the only correct answer in this case.
```

Constraints:

0 <= a, b, c <= 100  
a + b + c > 0


## 方法

```go
type item struct {
    name string
    num int
}

type myList []item

func (list myList) Len() int {
    return len(list)
}

func (list myList) Less(i int, j int) bool {
    return list[i].num > list[j].num
}

func (list myList) Swap(i int, j int) {
    list[i], list[j] = list[j], list[i]
}


func longestDiverseString(a int, b int, c int) string {
    list := myList{item{name:"a", num:a},
                   item{name:"b", num:b},
                   item{name:"c", num:c}}
    sort.Sort(list)
    
    res := ""
    
    for {
        if len(res) <= 1 || res[len(res)-1] != res[len(res)-2] || string(res[len(res)-1]) != list[0].name {
            if list[0].num <= 0 {
                break
            }
            res += list[0].name
            list[0].num--
            sort.Sort(list)
        } else {
            if list[1].num <= 0 {
                break
            }
            res += list[1].name
            list[1].num--
            sort.Sort(list)
        }
    }
    return res
}
```



```python
class Solution:
    def longestDiverseString(self, a: int, b: int, c: int) -> str:
        x = [[a, 'a'],[b, 'b'],[c,'c']]
        res = ""
        while True:
            for num in sorted(x,reverse=True):
                if num[0] <= 0:
                    return res
                if len(res) >= 2 and res[-2:] == num[1] * 2:
                    continue
                res += num[1]
                num[0] -= 1
                break
        return res
```