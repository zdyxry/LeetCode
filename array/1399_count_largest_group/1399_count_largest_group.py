class Solution:
    def countLargestGroup(self, n: int) -> int:
        d = {}
        def dsum(num):
            s = 0
            while num:
                s += num % 10
                num = num // 10
            return s
        
        for i in range(1, n+1):
            ds = dsum(i)
            if ds not in d:
                d[ds] = 1
            else:
                d[ds] += 1
        return len([v for v in d.values() if v == max(d.values())])


n = 13
res = Solution().countLargestGroup(n)
print(res)