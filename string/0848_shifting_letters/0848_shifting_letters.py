class Solution(object):
    def shiftingLetters(self, S, shifts):
        result = []
        times = sum(shifts) % 26
        for i, c in enumerate(S):
            index = ord(c) - ord('a')
            result.append(chr(ord('a') + (index+times) % 26))
            times = (times - shifts[i]) % 26
        return "".join(result)


S = "abc"
shifts = [3,5,9]
res = Solution().shiftingLetters(S, shifts)
print(res)