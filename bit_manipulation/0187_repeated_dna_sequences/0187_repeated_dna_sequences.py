class Solution(object):
    def findRepeatedDnaSequences(self, s):
        """
        :type s: str
        :rtype: List[str]
        """
        dict, rolling_hash, res = {}, 0, []

        for i in xrange(len(s)):
            rolling_hash = ((rolling_hash << 3) & 0x3fffffff) | (ord(s[i]) & 7)
            if rolling_hash not in dict:
                dict[rolling_hash] = True
            elif dict[rolling_hash]:
                res.append(s[i - 9: i + 1])
                dict[rolling_hash] = False
        return res

s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
res = Solution().findRepeatedDnaSequences(s)
print(res)