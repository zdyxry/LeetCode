class Solution:
    def countGoodSubstrings(self, s: str) -> int:
        start = 0
        end = len(s)
        result = []
        while start <= end -3:
            if len(s[start:start+3]) == len(set(s[start:start+3])):
                result.append(s[start:start+3])
            start += 1
        return len(result)
            

s = "xyzzaz"
print(Solution().countGoodSubstrings(s))