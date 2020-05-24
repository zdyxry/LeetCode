from typing import List

class Solution:
    def peopleIndexes(self, favoriteCompanies: List[List[str]]) -> List[int]:
        l = favoriteCompanies
        for i in range(len(l)):
            l[i] = [l[i], i]
        l.sort(key=lambda x: len(x[0]))
        print(l)
        ans = []
        for i in range(len(l)):
            for j in range(i+1, len(l)):
                if set(l[i][0]).issubset(set(l[j][0])):
                    break
            else:
                ans.append(l[i][1])
        ans.sort()
        return ans


favoriteCompanies = [["leetcode","google","facebook"],["google","microsoft"],["google","facebook"],["google"],["amazon"]]
res = Solution().peopleIndexes(favoriteCompanies)
print(res)