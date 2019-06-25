class Solution(object):
    def uncommonFromSentences(self, A, B):
        """
        :type A: str
        :type B: str
        :rtype: List[str]
        """
        res = []
        d = {}
        for s in A.split(" ") + B.split(" "):
            if s in d:
                d[s] += 1
            else:
                d[s] = 1
        
        for k,v in d.items():
            if v == 1:
                res.append(k)

        return res


A = "this apple is sweet"
B = "this apple is sour"

print(Solution().uncommonFromSentences(A, B))