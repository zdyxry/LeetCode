class Solution(object):
    def letterCombinations(self, digits):
        if not digits:
            return []
        
        letters = { "2": ["a","b","c"],
                    "3": ["d","e","f"],
                    "4": ["g","h","i"],
                    "5": ["j","k","l"],
                    "6": ["m","n","o"],
                    "7": ["p","q","r","s"],
                    "8": ["t","u","v"],
                    "9": ["w","x","y","z"]}

        retList = [""]
        for d in digits:
            retList = [s + c for s in retList for c in letters[d]]
        
        return retList

digits = "23"
res = Solution().letterCombinations(digits)
print(res)