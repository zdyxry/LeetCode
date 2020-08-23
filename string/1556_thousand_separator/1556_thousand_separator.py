class Solution:
    def thousandSeparator(self, n: int) -> str:
        n = str(n)[::-1]
        n1 = ""
        c = 0
        for i in n:
            c += 1
            n1 += i
            if c % 3 == 0:
                n1 += "."
        return n1[::-1].strip(".")


n = 987
res = Solution().thousandSeparator(n)
print(res)