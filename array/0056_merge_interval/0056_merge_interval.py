class Solution(object):
    def merge(self, intervals):
        if not intervals:
            return []

        length = len(intervals)
        if length == 1:
            return intervals
        intervals.sort(key=lambda x: x[0])
        res = [intervals[0]]

        for i in range(length):
            if intervals[i][0] <= res[-1][1]:
                res[-1][1] = max(intervals[i][1], res[-1][1])
            else:
                res.append(intervals[i])
        return res


intervals = [[1,3],[2,6],[8,10],[15,18]]
res = Solution().merge(intervals)
print(res)