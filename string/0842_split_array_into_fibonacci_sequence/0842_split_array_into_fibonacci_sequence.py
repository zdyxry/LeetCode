
class Solution(object):
    def splitIntoFibonacci(self, S):
        res = []
        current = []
        def backtrack(cursor, current):
            if len(current) >= 3 and cursor == len(S):
                res.append(current)
                return
            if len(current) < 2 :
                for index in range(cursor, len(S)):
                    avai=S[cursor:index+1]
                    if (avai[0] == "0" and len(avai)>1) or int(avai)>2**31-1:
                        return
                    backtrack(index+1, current+[int(avai)])
            else:
                a = current[-1]
                b = current[-2]
                seek = str(a+b)
                if S[cursor:cursor+len(seek)] == seek and int(seek)<2**31-1:
                    backtrack(cursor+len(seek), current+[int(seek)])
                else:
                    return
        backtrack(0, current)
        if res:
            return res[0]
        else:
            return res


S = "123456579" 
res = Solution().splitIntoFibonacci(S)
print(res)