class Solution:
    def makeEqual(self, words: List[str]) -> bool:
        cnt = [0] * 26
        n = len(words)
        for wd in words:
            for ch in wd:
                cnt[ord(ch)-ord('a')] += 1
        return all(k % n == 0 for k in cnt)
