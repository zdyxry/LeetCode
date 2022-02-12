class Solution:
    def checkAlmostEquivalent(self, word1: str, word2: str) -> bool:
        freq = defaultdict(int)   # 频数差哈希表
        for ch in word1:
            freq[ch] += 1
        for ch in word2:
            freq[ch] -= 1
        # 判断每个字符频数差是否均小于等于 3
        return all(abs(x) <= 3 for x in freq.values())