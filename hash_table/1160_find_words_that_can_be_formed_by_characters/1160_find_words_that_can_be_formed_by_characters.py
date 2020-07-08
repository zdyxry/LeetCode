from typing import List
import collections


class Solution:
    def countCharacters(self, words: List[str], chars: str) -> int:
        chars_cnt = collections.Counter(chars)
        ans = 0
        for word in words:
            word_cnt = collections.Counter(word)
            for c in word_cnt:
                if chars_cnt[c] < word_cnt[c]:
                    break
            else:
                ans += len(word)
        return ans


words = ["cat","bt","hat","tree"]
chars = "atach"
res = Solution().countCharacters(words, chars)
print(res)