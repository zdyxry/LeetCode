1154. Day of the Year


Easy


Given a string date representing a Gregorian calendar date formatted as YYYY-MM-DD, return the day number of the year.

 

Example 1:

```
Input: date = "2019-01-09"
Output: 9
Explanation: Given date is the 9th day of the year in 2019.
```

Example 2:

```
Input: date = "2019-02-10"
Output: 41
```

Example 3:

```
Input: date = "2003-03-01"
Output: 60
```

Example 4:

```
Input: date = "2004-03-01"
Output: 61
```
 

Constraints:

date.length == 10  
date[4] == date[7] == '-', and all other date[i]'s are digits  
date represents a calendar date between Jan 1st, 1900 and Dec 31, 2019.


## 方法

```go
func dayOfYear(date string) int {
    order := 0
	ds := strings.Split(date, "-")
	if len(ds) != 3 {
		return order
	}
	days := [12]int{31,28,31,30,31,30,31,31,30,31,30,31}
	y, _ := strconv.Atoi(ds[0])
	m, _ := strconv.Atoi(ds[1])
	d, _ := strconv.Atoi(ds[2])
	order += d
	for i := 0; i < m - 1; i++ {
		order += days[i]
	}
	if m > 2 && ((y % 100 == 0 && y % 400 == 0) || (y % 100 != 0 && y % 4 == 0)) {
		order++
	}
	return order
}
```


```python
class Solution(object):
    def dayOfYear(self, date):
        """
        :type date: str
        :rtype: int
        """
        y, m, d = map(int, date.split('-'))
        days = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
        if (y % 400) == 0 or ((y % 4 == 0) and (y % 100 != 0)): days[1] = 29
        return d + sum(days[:m-1])
```