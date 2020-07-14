class Solution:
    def numSub(self, s: str) -> int:
        cnt = 0
        res = 0
        for i in s:
            if i == '1':
                cnt += 1
            else:
                res += cnt * (cnt +1) /2
                cnt = 0
        if cnt > 0:
            res += cnt * (cnt + 1)/ 2
        res %= 1e9+7
        return int(res)
                
                

s = "0110111"
res = Solution().numSub(s)
print(res)