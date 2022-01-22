class Solution:
    def minMoves(self, target: int, maxDoubles: int) -> int:
        if maxDoubles == 0 :
            return target - 1

        result = 0
        while target >1 :
            if maxDoubles > 0:
                if target % 2 == 0:
                    result += 1
                    maxDoubles -= 1
                    target = target //2
                else:
                    result += 1
                    target -= 1
            else:
                result += target - 1
                target = 1
        return result
                    

