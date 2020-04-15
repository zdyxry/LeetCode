1138. Alphabet Board Path


Medium


On an alphabet board, we start at position (0, 0), corresponding to character board[0][0].

Here, board = ["abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"], as shown in the diagram below.



We may make the following moves:

```
'U' moves our position up one row, if the position exists on the board;
'D' moves our position down one row, if the position exists on the board;
'L' moves our position left one column, if the position exists on the board;
'R' moves our position right one column, if the position exists on the board;
'!' adds the character board[r][c] at our current position (r, c) to the answer.
(Here, the only positions that exist on the board are positions with letters on them.)
```

Return a sequence of moves that makes our answer equal to target in the minimum number of moves.  You may return any path that does so.

 

Example 1:

```
Input: target = "leet"
Output: "DDR!UURRR!!DDD!"
```

Example 2:

```
Input: target = "code"
Output: "RR!DDRR!UUL!R!"
```

Constraints:

1 <= target.length <= 100  
target consists only of English lowercase letters.


## 方法


```go
func alphabetBoardPath(target string) string {
    x, y := 0, 0
    moves := make([]byte, 0)
    
    for _, s := range target {
        n := int(s - 'a')
        nx, ny := n%5, n/5
        dx, dy := nx-x, ny-y
        
        if y == 5 {
            if dy != 0 {
                moves = append(moves, 'U')
                dy += 1
            }
        }
        
        for i:=0;i<abs(dx);i++ {
            if dx > 0 {
                moves = append(moves, 'R')
            } else {
                moves = append(moves, 'L')
            }
        }
        
        for i:=0;i<abs(dy);i++ {
             if dy > 0 {
                moves = append(moves, 'D')
            } else {
                moves = append(moves, 'U')
            }
        }
        
        moves = append(moves, '!')
        x, y = nx, ny
    }
    
    return string(moves)
}

func abs(a int) int {
    if a >= 0 {
        return a
    }
    return -a
}
```




```python
class Solution:
    def alphabetBoardPath(self, target: str) -> str:
        m = {c: [i // 5, i % 5] for i, c in enumerate("abcdefghijklmnopqrstuvwxyz")}
        x0, y0 = 0, 0
        res = []
        for c in target:
            x, y = m[c]
            if y < y0: res.append('L' * (y0 - y))
            if x < x0: res.append('U' * (x0 - x))
            if x > x0: res.append('D' * (x - x0))
            if y > y0: res.append('R' * (y - y0))
            res.append('!')
            x0, y0 = x, y
        return "".join(res)
```