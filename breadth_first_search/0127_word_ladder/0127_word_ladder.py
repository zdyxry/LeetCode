
class Solution(object):
    def ladderLength(self, beginWord, endWord, wordList):
        from string import ascii_lowercase
        distance, cur, visited, lookup = 0, [beginWord], set([beginWord]), set(wordList)

        while cur:
            next_queue = []

            for word in cur:
                if word == endWord:
                    return distance + 1
                for i in xrange(len(word)):
                    for j in ascii_lowercase:
                        candidate = word[:i] + j + word[i+1:]
                        if candidate not in visited and candidate in lookup:
                            next_queue.append(candidate)
                            visited.add(candidate)
            distance += 1
            cur = next_queue
        return 0 


beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log","cog"]

res = Solution().ladderLength(beginWord, endWord, wordList)
print(res)