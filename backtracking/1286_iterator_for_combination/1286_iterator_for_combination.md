1286. Iterator for Combination


Medium


Design an Iterator class, which has:

```
A constructor that takes a string characters of sorted distinct lowercase English letters and a number combinationLength as arguments.
A function next() that returns the next combination of length combinationLength in lexicographical order.
A function hasNext() that returns True if and only if there exists a next combination.
```

Example:

```
CombinationIterator iterator = new CombinationIterator("abc", 2); // creates the iterator.

iterator.next(); // returns "ab"
iterator.hasNext(); // returns true
iterator.next(); // returns "ac"
iterator.hasNext(); // returns true
iterator.next(); // returns "bc"
iterator.hasNext(); // returns false
```

Constraints:

```
1 <= combinationLength <= characters.length <= 15
There will be at most 10^4 function calls per test.
It's guaranteed that all calls of the function next are valid.
```


## 方法



```go

type CombinationIterator struct {
    index int
    combination []string
}

func Constructor(characters string, combinationLength int) CombinationIterator {
    combination := []string{}
    helper(0, combinationLength, "", characters, &combination)
    return CombinationIterator{ 0, combination }
}

func helper(start, length int, s, chars string, combination *[]string) {
    if length == 0 { *combination = append(*combination, s); return }
    for i := start; i < len(chars); i++ { helper(i + 1, length - 1, s + string(chars[i]), chars, combination) }
}

func (this *CombinationIterator) Next() string {
    tmp := this.combination[this.index]; this.index++; return tmp
}

func (this *CombinationIterator) HasNext() bool {
    return this.index < len(this.combination)
}
```


```python
class CombinationIterator:

    def __init__(self, characters: str, combinationLength: int):
        res=[]
        def helper(chars,length,path):
            if length==0:
                res.append(path[:])
            for index in range(len(chars)):
                path+=chars[index]
                helper(chars[index+1:],length-1,path)
                path=path[:-1]
        helper(characters,combinationLength,"")
        self.res=res
        self.res_len=len(res)
        
    def next(self) -> str:
        if self.hasNext():
            self.res_len-=1
            return self.res.pop(0)
        
    def hasNext(self) -> bool:
        if self.res_len==0:
            return False
        return True

# Your CombinationIterator object will be instantiated and called as such:
# obj = CombinationIterator(characters, combinationLength)
# param_1 = obj.next()
# param_2 = obj.hasNext()
```