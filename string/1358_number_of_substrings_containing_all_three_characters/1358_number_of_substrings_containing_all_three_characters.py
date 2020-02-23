
class Solution(object):
    def numberOfSubstrings(self, s):
        res = i = 0
        count = {c: 0 for c in 'abc'}
        for j in xrange(len(s)):
            count[s[j]] += 1
            while all(count.values()):
                count[s[i]] -= 1
                i += 1
            res += i
        return res

    def numberOfSubstrings2(self, s):
        ans, lo, cnt = 0, -1, {c:0 for c in 'abc'}
        for hi, c in enumerate(s):
            cnt[c] += 1
            while all(cnt.values()):
                ans += len(s) - hi
                lo += 1
                cnt[s[lo]] -= 1
            return ans


s = "abcc"
res = Solution().numberOfSubstrings2(s)
print(res)