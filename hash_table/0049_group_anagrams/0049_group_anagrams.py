import collections

class Solution(object):
    def groupAnagrams(self, strs):
        anagrams_map, result = collections.defaultdict(list), []
        for s in strs:
            sorted_str = ("").join(sorted(s))
            anagrams_map[sorted_str].append(s)

        for anagram in anagrams_map.values():
            anagram.sort()
            result.append(anagram)

        return result

strs = ["eat","tea","tan","ate","nat","bat"]
res = Solution().groupAnagrams(strs)
print(res)