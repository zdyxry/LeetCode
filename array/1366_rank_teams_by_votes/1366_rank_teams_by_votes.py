
import collections

class Solution(object):
    def rankTeams(self, votes):
        n = len(votes[0])
        ranking = collections.defaultdict(lambda : [0] * n)
        for vote in votes:
            for i, vid in enumerate(vote):
                ranking[vid][i] += 1
        
        result = list(ranking.items())
        result.sort(key=lambda x: (x[1], -ord(x[0])),reverse=True)
        return "".join([vid for vid, _ in result])


votes =["ABC","ACB","ABC","ACB","ACB"]
res = Solution().rankTeams(votes)
print(res)