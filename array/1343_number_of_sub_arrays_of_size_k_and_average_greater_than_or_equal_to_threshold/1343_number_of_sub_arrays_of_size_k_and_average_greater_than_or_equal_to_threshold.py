
class Solution(object):
    def numOfSubarrays(self, arr, k, threshold):
        n = len(arr)
        left = 0
        res = 0
        for i in range(n):
            left += arr[i]
            if i >= k-1:
                if left >= threshold * k:
                    res += 1
                left -= arr[i-k+1]
        return res


arr = [2,2,2,2,5,5,5,8]
k = 3
threshold = 4

res = Solution().numOfSubarrays(arr, k, threshold)
print(res)