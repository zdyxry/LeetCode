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


day = 31
month = 8
year = 2019
res = Solution().dayOfTheWeek(day, month, year)
print(res)