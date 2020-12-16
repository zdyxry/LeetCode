class Solution:
    def minPartitions(self, n: str) -> int:
        return (int(max(list(str(n)))))


n = 32
res = Solution().minPartitions(n)
print(res)