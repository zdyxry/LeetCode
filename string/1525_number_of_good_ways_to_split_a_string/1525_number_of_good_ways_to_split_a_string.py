from collections import Counter


class Solution:
    def numSplits(self, s: str) -> int:
        left = Counter()
        right = Counter(s)
        count = 0
        for i in range(len(s)):
            left[s[i]] += 1
            right[s[i]] -= 1
            if right[s[i]] == 0:
                del right[s[i]]
                
            if len(left) == len(right):
                count += 1
        return count


s = "aacaba"
res = Solution().numSplits(s)
print(res)