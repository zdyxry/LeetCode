class Solution:
    def checkZeroOnes(self, s: str) -> bool:
        i = 0 
        cnt0 = 0
        cnt1 = 0
        result0 = 0
        result1 = 0
        while i < len(s):
            if s[i] == '0':
                cnt0 += 1
                cnt1 = 0
            if s[i] == '1':
                cnt1 += 1
                cnt0 = 0
            i += 1
            result0 = max(result0, cnt0)
            result1 = max(result1, cnt1)
        return result1 > result0


s = "1101"
res = Solution().checkZeroOnes(s)
print(res)