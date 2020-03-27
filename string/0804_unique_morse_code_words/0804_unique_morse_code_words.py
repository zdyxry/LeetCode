
class Solution(object):
    def uniqueMorseRepresentations(self, words):
        morse_alphabet = [".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",
                          ".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",
                          ".--","-..-","-.--","--.."]
        transformations = set()
        for word in words:
            morse_str = "".join([morse_alphabet[ord(c)-ord("a")]for c in word])
            transformations.add(morse_str)
        return len(transformations)


words = ["gin", "zen", "gig", "msg"]
res = Solution().uniqueMorseRepresentations(words)
print(res)