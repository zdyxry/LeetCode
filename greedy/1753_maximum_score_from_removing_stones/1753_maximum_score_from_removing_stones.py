class Solution:
    def maximumScore(self, a: int, b: int, c: int) -> int:
        a,b,c = sorted([a,b,c])
        if a+b <= c:
            return (a+b)
        else:
            return (a+b+c)//2


a, b, c = 2,4,6
res = Solution().maximumScore(a,b,c)
print(res)