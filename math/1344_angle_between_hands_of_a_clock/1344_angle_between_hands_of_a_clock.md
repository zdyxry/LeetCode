1344. Angle Between Hands of a Clock


Medium


Given two numbers, hour and minutes. Return the smaller angle (in sexagesimal units) formed between the hour and the minute hand.

 

Example 1:

```
Input: hour = 12, minutes = 30
Output: 165
```

Example 2:

```
Input: hour = 3, minutes = 30
Output: 75
```

Example 3:

```
Input: hour = 3, minutes = 15
Output: 7.5
```

Example 4:

```
Input: hour = 4, minutes = 50
Output: 155
```

Example 5:

```
Input: hour = 12, minutes = 0
Output: 0
```

 

Constraints:

1 <= hour <= 12  
0 <= minutes <= 59  
Answers within 10^-5 of the actual value will be accepted as correct.

## 方法

```go
func angleClock(hour int, minutes int) float64 {
    var (
        locaH, locaM, diff float64
    )

    locaM = float64(6 * minutes)
    locaH = float64(30*hour%360) + float64(minutes)*30.0/60.0

    diff = max(locaH, locaM) - min(locaH, locaM)
    if diff > 180.0 {
        diff = 360 - diff
    }
    return diff
}

func max(a, b float64) float64 {
    if a > b {
        return a
    }
    return b
}
func min(a, b float64) float64 {
    if a < b {
        return a
    }
    return b
}
```


```python
class Solution(object):
    def angleClock(self, hour, minutes):
        """
        :type hour: int
        :type minutes: int
        :rtype: float
        """
        m = minutes * (360/60)
        if hour == 12:
            h = minutes * (360.0/12/60)
        else:
            h = hour * (360/12) + minutes * (360.0/12/60)
        res = abs(m-h)
        return 360 - res if res > 180 else res
```