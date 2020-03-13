
class Solution(object):
    def dayOfYear(self, date):
        y, m, d = map(int, date.split('-'))
        days = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
        if (y % 400) == 0 or ((y % 4 ==0 ) and (y % 100 != 0)):
            days[1] = 29
        return d + sum(days[:m-1])


date = "2019-01-09"
res = Solution().dayOfYear(date)
print(res)