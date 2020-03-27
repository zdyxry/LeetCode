804. Unique Morse Code Words


Easy


International Morse Code defines a standard encoding where each letter is mapped to a series of dots and dashes, as follows: "a" maps to ".-", "b" maps to "-...", "c" maps to "-.-.", and so on.

For convenience, the full table for the 26 letters of the English alphabet is given below:

```
[".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]
Now, given a list of words, each word can be written as a concatenation of the Morse code of each letter. For example, "cba" can be written as "-.-..--...", (which is the concatenation "-.-." + "-..." + ".-"). We'll call such a concatenation, the transformation of a word.
```

Return the number of different transformations among all words we have.

Example:
```
Input: words = ["gin", "zen", "gig", "msg"]
Output: 2
Explanation: 
The transformation of each word is:
"gin" -> "--...-."
"zen" -> "--...-."
"gig" -> "--...--."
"msg" -> "--...--."

There are 2 different transformations, "--...-." and "--...--.".
```

Note:

The length of words will be at most 100.  
Each words[i] will have length in range [1, 12].  
words[i] will only consist of lowercase letters.


## 方法

```go
func uniqueMorseRepresentations(words []string) int {
    morses := [] string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
    transformations := make(map[string]bool)
    for _, s := range words {
        transformation := ""
        runes := [] rune(s)
        for _, r := range runes {
            index := r - 'a'
            transformation += morses[index]
        }
        transformations[transformation] = true
    }
    return len(transformations)
}
```



```python
class Solution(object):
    def uniqueMorseRepresentations(self, words):
        """
        :type words: List[str]
        :rtype: int
        """
        morse_alphabet = [".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",
                          ".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",
                          ".--","-..-","-.--","--.."]
        transformations = set()
        for word in words:
            morse_str = ''.join([morse_alphabet[ord(c) - ord('a')] for c in word])
            transformations.add(morse_str)
        return len(transformations)
```