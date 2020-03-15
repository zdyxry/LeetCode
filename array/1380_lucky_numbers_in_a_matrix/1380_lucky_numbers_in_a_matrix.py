
class Solution(object):
    def luckyNumbers(self, matrix):
        mins = {min(rows) for rows in matrix}
        maxes = {max(columns) for columns in zip(*matrix)}
        return list(mins & maxes)


matrix = [[3,7,8],[9,11,13],[15,16,17]]
res = Solution().luckyNumbers(matrix)
print(res)