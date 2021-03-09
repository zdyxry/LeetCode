class Solution:
    def checkOnesSegment(self, s: str) -> bool:
        new = 0
        for i in s:
            if i == '0':
                new = 1
            else:
                if new == 1:
                    return False
        return True


s = '1001'
res = Solution().checkOnesSegment(s)
print(res)