
class Solution:
    def judgeCircle(self, moves: str) -> bool:
        return moves.count('L') == moves.count('R') and moves.count('U') == moves.count('D')

moves = "UD"
res = Solution().judgeCircle(moves)
print(res)