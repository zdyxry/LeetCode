
class Solution:
    def minNumberOfFrogs(self, croakOfFrogs: str) -> int:
        c, r, o, a, k, in_use, answer = 0, 0, 0, 0, 0, 0, 0
        for d in croakOfFrogs:
            if d == 'c':
                c, in_use = c+1, in_use+1
            elif d == 'r':
                r += 1
            elif d == 'o':
                o += 1
            elif d == 'a':
                a += 1
            else:
                k, in_use = k+1, in_use - 1
            
            answer = max(answer, in_use)

            if c < r or r < o or o < a or a < k:
                return -1

        if in_use == 0 and c == r and r == o and o == a and a == k:
            return answer
    
        return -1


croakOfFrogs = "croakcroak"
res = Solution().minNumberOfFrogs(croakOfFrogs)
print(res)