class Solution:
    def isPathCrossing(self, path: str) -> bool:
        x = y = 0
        set = {(0, 0)}
        
        for c in path:
            if c == 'N':
                y += 1
            elif c == 'S':
                y -= 1
            elif c == 'E':
                x += 1
            else:
                x -= 1
            
            if (x, y) in set:
                return True
            set.add((x, y))
        return False


path = "NES"
res = Solution().isPathCrossing(path)
print(res)