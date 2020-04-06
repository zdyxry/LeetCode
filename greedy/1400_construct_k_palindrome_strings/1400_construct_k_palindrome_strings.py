
class Solution:
    def canConstruct(self, s: str, k: int) -> bool:
        digit_count = {}
        if len(s) < k:
            return False
        elif len(s) == k:
            return True
        else:
            odd = 0
            for i in set(s):
                digit_count[i] = s.count(i)
            
            for i in digit_count.values():
                if i % 2 != 0:
                    odd += 1
            if odd > k:
                return False
            else:
                return True


s = "annabelle"
k = 2
res = Solution().canConstruct(s, k)
print(res)