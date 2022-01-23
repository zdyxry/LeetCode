class Solution:
    def executeInstructions(self, n: int, startPos: List[int], s: str) -> List[int]:
        matrix = [[0] * n] * n 
        result = []
        for i in range(0, len(s)):
            r, c = startPos[0], startPos[1]
            matrix[r][c] = 1
            tmp = 0
            for cmd in s[i:]:
                if cmd == 'U':
                    if r > 0:
                        r -= 1
                        tmp += 1
                    else:
                        break
                elif cmd == 'D':
                    if r < n-1:
                        r += 1
                        tmp += 1
                    else:
                        break
                elif cmd == 'L':
                    if c > 0:
                        c -= 1
                        tmp += 1
                    else:
                        break
                elif cmd == 'R':
                    if c < n-1:
                        c += 1
                        tmp += 1
                    else:
                        break
            result.append(tmp)
        return result
                

