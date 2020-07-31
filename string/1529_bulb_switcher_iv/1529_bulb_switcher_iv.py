class Solution:
    def minFlips(self, target: str) -> int:
        flips = 0
        i0 = target.find('1')
        if i0 ==-1:
            return flips
        flips = 1
        for i in range(i0 + 1, len(target)):
            if target[i] != target[i-1]:
                flips += 1
            
        return flips


target = "10111"
res = Solution().minFlips(target)
print(res)