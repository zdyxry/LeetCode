class Solution(object):
    def uniquePaths(self, m, n):
        if m < n:
            return self.uniquePaths(n, m)
        ways = [1] * n

        for i in xrange(1, m):
            for j in xrange(1, n):
                print("#####")
                print(ways)
                ways[j] += ways[j - 1]
                print(ways)
        return ways[n - 1]


m = 4
n = 3
res = Solution().uniquePaths(m, n)
print(res)