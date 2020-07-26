1185. Day of the Week


Easy


Given a date, return the corresponding day of the week for that date.

The input is given as three integers representing the day, month and year respectively.

Return the answer as one of the following values {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}.

 

Example 1:

```
Input: day = 31, month = 8, year = 2019
Output: "Saturday"
```

Example 2:

```
Input: day = 18, month = 7, year = 1999
Output: "Sunday"
```

Example 3:

```
Input: day = 15, month = 8, year = 1993
Output: "Sunday"
```
 

Constraints:

The given dates are valid dates between the years 1971 and 2100.


## 方法

```go

func isLeapYear(year int) bool{
	if (year%4==0 && year%100!=0) || (year%400==0) {
		return true
	}
	return  false
}

func getDay(year int, mon int, day int) (res int){
	months := []int{-1, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for i := 1970; i < year; i++{
		res += 365
		if isLeapYear(i){
			res++
		}
	}
	for i := 1; i < mon; i++{
		res += months[i]
		if i == 2 && isLeapYear(year) {
			res++
		}
	}
	res += day
	return res
}

func dayOfTheWeek(day int, month int, year int) string {
	weekMap := map[int]string{
		1: "Thursday",
		2: "Friday",
		3: "Saturday",
		4: "Sunday",
		5: "Monday",
		6: "Tuesday",
		0: "Wednesday",
	}
	d := getDay(year, month, day)
	return weekMap[d%7]
}

```



```python
class Solution:
    def dayOfTheWeek(self, day: int, month: int, year: int) -> str:
        total = 0
        monthDateN = [31,28,31,30,31,30,31,31,30,31,30,31]
        monthDateL = [31,29,31,30,31,30,31,31,30,31,30,31]
        for i in range(1970, year):
            if self.isLeapYear(i):
                total += 366
            else:
                total += 365
        if self.isLeapYear(year):
            for i in range(1, month):
                total += monthDateL[i - 1]
        else:
            for i in range(1, month):
                total += monthDateN[i - 1]
        total += day
        weekday = ["Wednesday","Thursday", "Friday", "Saturday","Sunday", "Monday", "Tuesday"]
        return weekday[total % 7]
            
    def isLeapYear(self, year):
        return (year % 4 == 0 and year % 100 != 0) or year % 400 == 0
```