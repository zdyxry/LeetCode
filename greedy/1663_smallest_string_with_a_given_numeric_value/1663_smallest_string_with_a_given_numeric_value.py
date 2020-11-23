class Solution:
    def getSmallestString(self, n: int, k: int) -> str:
        k -= n
        cnt_z = k // 25
        last = k % 25
        return 'a' * (n - 1 - cnt_z) + chr(ord('a') + last) + 'z' * cnt_z


n = 3
k = 27
res = Solution().getSmallestString(n, k)
print(res)
