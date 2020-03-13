
class Solution(object):
    def numOfMinutes(self, n, headID, manager, informTime):
        def dfs(i):
            if manager[i] != -1:
                informTime[i] += dfs(manager[i])
                manager[i] = -1
            return informTime[i]
        return max(map(dfs, range(n)))


n = 1
headID = 0
manager = [-1]
informTime = [0]

res = Solution().numOfMinutes(n, headID, manager, informTime)
print(res)