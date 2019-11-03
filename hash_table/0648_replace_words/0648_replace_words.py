import collections
from pprint import pprint


class Solution(object):
    def replaceWords(self, dict, sentence):
        """
        :type dict: List[str]
        :type sentence: str
        :rtype: str
        """
        rootset = set(dict)
        
        def replace(word):
            for i in xrange(1, len(word)):
                if word[:i] in rootset:
                    return word[:i]
            return word
        return " ".join(map(replace, sentence.split()))

    def replaceWords2(self, dict, sentence):
        _trie = lambda : collections.defaultdict(_trie)
        trie = _trie()
        END = True

        for root in dict:
            cur = trie
            for letter in root:
                cur = cur[letter]
            cur[END] = root
        
        pprint(trie)
        def replace(word):
            cur = trie
            for letter in word:
                if letter not in cur:break
                cur = cur[letter]
                if END in cur:
                    return cur[END]
            return word
        return " ".join(map(replace, sentence.split()))



dict = ["cat", "bat", "rat"]
sentence = "the cattle was rattled by the battery"

res = Solution().replaceWords2(dict, sentence)
print(res)