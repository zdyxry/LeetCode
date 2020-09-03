728. Self Dividing Numbers


Easy

A self-dividing number is a number that is divisible by every digit it contains.

For example, 128 is a self-dividing number because 128 % 1 == 0, 128 % 2 == 0, and 128 % 8 == 0.

Also, a self-dividing number is not allowed to contain the digit zero.

Given a lower and upper number bound, output a list of every possible self dividing number, including the bounds if possible.

Example 1:
```
Input: 
left = 1, right = 22
Output: [1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]
```

Note:

The boundaries of each input argument are 1 <= left <= right <= 10000.


## 方法



```go
func isDividingNumbers(i int)bool{
	if i<10{
		return true
	}
	v:=i
	for i!=0{
		r :=i%10
		if r==0||v%r!=0{
			return false
		}
		i =i/10
	}
	return true
}
func selfDividingNumbers(left int, right int) []int {
	 var res []int
     for i:=left;i<=right;i++{
		 if isDividingNumbers(i){
		 	res = append(res,i)
		 }
	 }
	 return res
}


```


```python
class Solution:
    def selfDividingNumbers(self, left: int, right: int) -> List[int]:
        res = []
        for i in range(left, right+1):
            for j in list(str(i)):
                if int(j) == 0:
                    break
                if i % int(j) != 0:
                    break
            else:
                res.append(i)
        return res
```