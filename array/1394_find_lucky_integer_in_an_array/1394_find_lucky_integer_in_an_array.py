
class Solution(object):
    def findLucky(self, arr):
        cnt = [0] * 501
        for a in arr:
            cnt[a] += 1
        for i in range(500, 0, -1):
            if cnt[i] == i:
                return i
        return -1


arr = [2,2,3,4]
res = Solution().findLucky(arr)
print(res)