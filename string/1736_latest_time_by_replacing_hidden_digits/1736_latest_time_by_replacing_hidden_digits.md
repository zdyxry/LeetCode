1736. Latest Time by Replacing Hidden Digits


Easy


You are given a string time in the form of hh:mm, where some of the digits in the string are hidden (represented by ?).

The valid times are those inclusively between 00:00 and 23:59.

Return the latest valid time you can get from time by replacing the hidden digits.

 

Example 1:

```
Input: time = "2?:?0"
Output: "23:50"
Explanation: The latest hour beginning with the digit '2' is 23 and the latest minute ending with the digit '0' is 50.
```

Example 2:

```
Input: time = "0?:3?"
Output: "09:39"
```

Example 3:

```
Input: time = "1?:22"
Output: "19:22"
```
 

Constraints:

time is in the format hh:mm.   
It is guaranteed that you can produce a valid time from the given string.


## 方法

```go
func maximumTime(time string) string {
     times:=[]byte(time)
     if times[0]=='?'{
     	if times[1]=='?'||times[1]<'4'{
			times[0]='2'
		}else {
			times[0]='1'
		}
	 }
	if times[1]=='?'{
		if times[0]=='2'{
			times[1]='3'
		}else {
			times[1]='9'
		}
	}
	if times[3]=='?'{
		times[3]='5'
	}
	if times[4]=='?'{
		times[4]='9'
	}
	return string(times)
}

```


```python
class Solution:
    def maximumTime(self, time: str) -> str:
        hh = time[:2]
        if hh[0] == '?' and hh[1] == '?':
            hh = '23'
        elif hh[0] == '?':
            if hh[1] < '4':
                hh = f'2{hh[1]}'
            else:
                hh = f'1{hh[1]}' 
        elif hh[1] == '?':
            if hh[0] == '2':
                hh = f'{hh[0]}3'
            else:
                hh = f'{hh[0]}9'
        mm = time[3:]
        if mm[0] == '?':
            mm = f'5{mm[1]}'
        if mm[1] == '?':
            mm = f'{mm[0]}9'
        return f'{hh}:{mm}'
```