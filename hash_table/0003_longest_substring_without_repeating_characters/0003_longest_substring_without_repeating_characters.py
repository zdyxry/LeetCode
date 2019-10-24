class Solution(object):
    def lengthOfLongestSubstring(self, s):
        longest, start, visited = 0, 0, [False for _ in xrange(256)]
        for i, char in enumerate(s):
            if visited[ord(char)]:
                while char != s[start]:
                    visited[ord(s[start])] = False
                    start += 1
                start += 1
            else:
                visited[ord(char)] = True
            longest = max(longest, i - start + 1)
        return longest

s = "abcabcbb"
res = Solution().lengthOfLongestSubstring(s)
print(res)