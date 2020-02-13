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


hour = 12
minutes = 30
res = Solution().angleClock(hour, minutes)
print(res)