class Solution:
    def checkString(self, s: str) -> bool:
        n = len(s)
        lasta = -1   # 'a' 最后一次出现的下标
        firstb = n   # 'b' 第一次出现的下标
        for i in range(n):
            if s[i] == 'a':
                lasta = max(lasta, i)
            else:
                firstb = min(firstb, i)
        return lasta < firstb
