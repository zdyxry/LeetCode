401. Binary Watch
Easy

A binary watch has 4 LEDs on the top which represent the hours (0-11), and the 6 LEDs on the bottom represent the minutes (0-59).

Each LED represents a zero or one, with the least significant bit on the right.
0
For example, the above binary watch reads "3:25".

Given a non-negative integer n which represents the number of LEDs that are currently on, return all possible times the watch could represent.

Example:

Input: n = 1  
Return: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]  
Note:  
The order of output does not matter.  
The hour must not contain a leading zero, for example "01:00" is not valid, it should be "1:00".  
The minute must be consist of two digits and may contain a leading zero, for example "10:2" is not valid, it should be "10:02".  

# 方法
遍历所有时间数字，如果当前时间数字二进制之和中 1 的个数等于 num，则记录。


```go
func readBinaryWatch(n int) []string {
	res := make([]string, 0, 8)
	leds := make([]bool, 10)

	var dfs func(int, int)
	dfs = func(idx, n int) {
		var h, m int
		if n == 0 {
			m, h = get(leds[:6]), get(leds[6:])
			if h < 12 && m < 60 {
				res = append(res, fmt.Sprintf("%d:%02d", h, m))
			}
			return
		}

		for i := idx; i < len(leds)-n+1; i++ {
			leds[i] = true
			dfs(i+1, n-1)
			leds[i] = false
		}
	}

	dfs(0, n)

	sort.Strings(res)

	return res
}

var bs = []int{1, 2, 4, 8, 16, 32}

func get(leds []bool) int {
	var sum int
	size := len(leds)
	for i := 0; i < size; i++ {
		if leds[i] {
			sum += bs[i]
		}
	}

	return sum
}
```


```python
def readBinaryWatch(self, num):
    result = []
    for h in range(12):
        for m in range(60):
            if (bin(h) + bin(m)).count('1') == num:
                result.append('%d:%02d' % (h, m))
    return result
```