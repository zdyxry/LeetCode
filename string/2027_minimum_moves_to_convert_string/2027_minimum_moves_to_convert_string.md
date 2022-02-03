2027. Minimum Moves to Convert String


Easy


You are given a string s consisting of n characters which are either 'X' or 'O'.

A move is defined as selecting three consecutive characters of s and converting them to 'O'. Note that if a move is applied to the character 'O', it will stay the same.

Return the minimum number of moves required so that all the characters of s are converted to 'O'.

 

Example 1:

```
Input: s = "XXX"
Output: 1
Explanation: XXX -> OOO
We select all the 3 characters and convert them in one move.
```

Example 2:

```
Input: s = "XXOX"
Output: 2
Explanation: XXOX -> OOOX -> OOOO
We select the first 3 characters in the first move, and convert them to 'O'.
Then we select the last 3 characters and convert them so that the final string contains all 'O's.
```

Example 3:

```
Input: s = "OOOO"
Output: 0
Explanation: There are no 'X's in s to convert.
```

Constraints:

3 <= s.length <= 1000   
s[i] is either 'X' or 'O'.


## 方法


```
func minimumMoves(s string) int {
    result := 0
    for i := 0 ; i < len(s) ; {
        if s[i] == 'X'{
            result += 1
            i += 3
        } else {
            i += 1
        }
    }
    return result
}
```


```
class Solution:
    def minimumMoves(self, s: str) -> int:
        l = len(s)
        i = 0
        result = 0
        while i < l:
            if s[i] == 'X':
                result += 1
                i += 3
            else:
                i += 1
        return result

```