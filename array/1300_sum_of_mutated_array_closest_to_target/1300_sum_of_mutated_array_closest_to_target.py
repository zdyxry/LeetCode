
class Solution(object):
    def findBestValue(self, arr, target):
        arr.sort(reverse = True)
        while arr and target >= arr[-1]*len(arr):
            temp = arr[-1]
            target -= arr.pop()

        if not arr:
            return temp
        res = target / float(len(arr))

        if res % 1 > 0.5:
            return int(res) + 1
        else:
            return int(res)


arr = [2,2,2]
target = 10
res = Solution().findBestValue(arr, target)
print(res)