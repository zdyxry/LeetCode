class Solution:
    def getLucky(self, s: str, k: int) -> int:
        before = ""
        for c in s:
            before += str(ord(c)-ord('a')+1)
        result = 0 
        for i in range(k):
            before = str(sum([int(c) for c in before]))
        return int(before)