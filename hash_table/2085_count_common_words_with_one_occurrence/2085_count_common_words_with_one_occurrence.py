class Solution:
    def countWords(self, words1: List[str], words2: List[str]) -> int:
        freq1 = Counter(words1)   # words1 中字符串的出现次数
        freq2 = Counter(words2)   # words2 中字符串的出现次数
        res = 0   # 出现过一次的公共字符串个数
        for word1 in freq1.keys():
            # 遍历 words1 出现的字符并判断是否满足要求
            if freq1[word1] == 1 and freq2[word1] == 1:
                res += 1
        return res
