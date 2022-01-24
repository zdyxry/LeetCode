class Solution:
    def minTimeToType(self, word: str) -> int:
        cur = 'a'
        result = 0
        for c in word:
            if c != cur:
                max_time = ord(c) - ord(cur)
                if max_time < 0:
                    max_time = abs(max_time)
                result += min(max_time, 26-max_time)
                result += 1
            else:
                result += 1
            cur = c
        return result
