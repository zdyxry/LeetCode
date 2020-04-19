import re
import itertools

class Solution(object):
    def reformat(self, s: str) -> str:
        a = re.findall(r'\d', s)
        b = re.findall(r'[a-z]', s)
        if abs(len(a) - len(b)) > 1:
            return ''
        a, b = sorted([a, b], key=len)
        return ''.join(map(''.join, itertools.zip_longest(a,b, fillvalue='')))


s = "a0b1c2"
res = Solution().reformat(s)
print(res)