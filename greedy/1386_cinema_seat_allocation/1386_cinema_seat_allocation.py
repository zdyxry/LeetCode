import collections


class Solution(object):
    def maxNumberOfFamilies(self, n, reservedSeats):
        seats = collections.defaultdict(int)
        res = 0

        for row, col in reservedSeats:
            seats[row] = seats[row] | (1 << (col-1))

        for reserved in seats.values():
            curr = 0
            curr += (reserved & int('0111100000', 2)) == 0
            curr += (reserved & int('0000011110', 2)) == 0
            curr += (reserved & int('0001111000', 2)) == 0 and curr == 0

            res += curr

        return res + 2 * (n- len(seats))



n = 3
reservedSeats = [[1,2],[1,3],[1,8],[2,6],[3,1],[3,10]]
res = Solution().maxNumberOfFamilies(n, reservedSeats)
print(res)