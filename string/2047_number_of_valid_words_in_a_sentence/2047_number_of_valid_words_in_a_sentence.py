class Solution:
    def countValidWords(self, sentence: str) -> int:
        result = 0
        for w in sentence.split():
            tmp = False
            for idx,c in enumerate(w):
                if str(c).isdigit():
                    break
                if c in '!.,' and idx != len(w)-1:
                        break
                if c == '-':
                    if tmp:
                        break
                    if idx == 0 or idx == len(w)-1:
                        break
                    if not w[idx-1].islower() or not w[idx+1].islower():
                        break
                    tmp = True
            else:
                result += 1
                tmp = False
        return result
