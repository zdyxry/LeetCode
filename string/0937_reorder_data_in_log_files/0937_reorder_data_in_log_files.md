937. Reorder Data in Log Files


Easy


You have an array of logs.  Each log is a space delimited string of words.

For each log, the first word in each log is an alphanumeric identifier.  Then, either:

```
Each word after the identifier will consist only of lowercase letters, or;
Each word after the identifier will consist only of digits.
We will call these two varieties of logs letter-logs and digit-logs.  It is guaranteed that each log has at least one word after its identifier.
```

Reorder the logs so that all of the letter-logs come before any digit-log.  The letter-logs are ordered lexicographically ignoring identifier, with the identifier used in case of ties.  The digit-logs should be put in their original order.

Return the final order of the logs.

 

Example 1:
```
Input: logs = ["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"]
Output: ["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"]
```

Constraints:

0 <= logs.length <= 100   
3 <= logs[i].length <= 100   
logs[i] is guaranteed to have an identifier, and a word after the identifier.


## 方法


```go
func reorderLogFiles(logs []string) []string {
    lettersMap := make(map[string]string)
	var letters []string
	var digits []string
   
    for _, log := range logs {
        v := strings.SplitN(log, " ", 2)
        if isDigit(v[1]) {
            digits = append(digits, log)
            continue
        } else {
            lettersMap[v[1]+v[0]] = log
        }
    }
    for k, _ := range lettersMap {
        letters = append(letters, k)
    }
    sort.Strings(letters)
    var result []string
    for _, l := range letters {
        result = append(result, lettersMap[l])
    }
    result = append(result, digits...)
    return result
}

func isDigit(log string) bool {
    for _, v := range log {
        vv := string(v)  
        _, err := strconv.Atoi(vv)
        if err != nil && vv != " " {
            return false
        }         
    }
    return true
}
```


```python
class Solution:
    def reorderLogFiles(self, logs: List[str]) -> List[str]:
        a=[]
        b=[]
        for i in logs:
            if i.split(" ")[1].isdigit():
                b.append(i)
            else:
                a.append(i)
        a=sorted(a,key=lambda a:(a.split(" ")[1:],a.split()[0]))
        return a+b
```