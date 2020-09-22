
class Solution:
    def maxUniqueSplit(self, s: str) -> int:
        self.ans = 0
        d = set()

        def helper(num, sub):
            if not sub:
                self.ans = max(num, self.ans)
                return
            
            if sub in d: return 
            
            for i in range(1, len(sub) + 1):
                if sub[:i] not in d:
                    d.add(sub[:i])
                    helper(num + 1, sub[i:])
                    d.remove(sub[:i]) 
                    
        helper(0, s)

        return self.ans



s = "ababccc"
res = Solution().maxUniqueSplit(s)
print(res)