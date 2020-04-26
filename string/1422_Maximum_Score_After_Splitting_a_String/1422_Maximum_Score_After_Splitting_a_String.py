class Solution:
    def maxScore(self, s: str) -> int:
        right = s.count('1')
        left = 0
        score = 0
        for idx in range(len(s) - 1):
            if s[idx] == '1':
                score = max(score, (left + right -1))
                right -= 1
            else:
                score = max(score, (left + 1 + right))
                left += 1
        return score