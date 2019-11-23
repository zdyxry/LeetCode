
class Solution(object):
    def wordSubsets(self, A, B):
        s = set(A)
        letters_required = {}
        for i in B:
            for j in i:
                count = i.count(j)
                if j not in letters_required or count > letters_required[j]:
                    letters_required[j] = count
        
        for i in A:
            for j in letters_required:
                if i.count(j) < letters_required[j]:
                    s.remove(i)
                    break
        return list(s)


A = ["amazon","apple","facebook","google","leetcode"]
B = ["e","o"]
res = Solution().wordSubsets(A, B)
print(res)