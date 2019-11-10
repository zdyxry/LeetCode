import collections


class Solution(object):
    def leastBricks(self, wall):
        widths = collections.defaultdict(int)
        result = len(wall)
        for row in wall:
            width = 0
            for i in xrange(len(row) - 1):
                width += row[i]
                widths[width] += i
                result = min(result, len(wall) - widths[width])
        return result


wall = [[1,2,2,1],[3,1,2],[1,3,2],[2,4],[3,1,2],[1,3,1,1]]

res = Solution().leastBricks(wall)
print(res)