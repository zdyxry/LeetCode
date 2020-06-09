1472. Design Browser History


Medium


You have a browser of one tab where you start on the homepage and you can visit another url, get back in the history number of steps or move forward in the history number of steps.

Implement the BrowserHistory class:

* BrowserHistory(string homepage) Initializes the object with the homepage of the browser.
* void visit(string url) visits url from the current page. It clears up all the forward history.
* string back(int steps) Move steps back in history. If you can only return x steps in the history and steps > x, you will return only x steps. Return the current url after moving back in history at most steps.
* string forward(int steps) Move steps forward in history. If you can only forward x steps in the history and steps > x, you will forward only x steps. Return the current url after forwarding in history at most steps.
 

Example:

```
Input:
["BrowserHistory","visit","visit","visit","back","back","forward","visit","forward","back","back"]
[["leetcode.com"],["google.com"],["facebook.com"],["youtube.com"],[1],[1],[1],["linkedin.com"],[2],[2],[7]]
Output:
[null,null,null,null,"facebook.com","google.com","facebook.com",null,"linkedin.com","google.com","leetcode.com"]

Explanation:
BrowserHistory browserHistory = new BrowserHistory("leetcode.com");
browserHistory.visit("google.com");       // You are in "leetcode.com". Visit "google.com"
browserHistory.visit("facebook.com");     // You are in "google.com". Visit "facebook.com"
browserHistory.visit("youtube.com");      // You are in "facebook.com". Visit "youtube.com"
browserHistory.back(1);                   // You are in "youtube.com", move back to "facebook.com" return "facebook.com"
browserHistory.back(1);                   // You are in "facebook.com", move back to "google.com" return "google.com"
browserHistory.forward(1);                // You are in "google.com", move forward to "facebook.com" return "facebook.com"
browserHistory.visit("linkedin.com");     // You are in "facebook.com". Visit "linkedin.com"
browserHistory.forward(2);                // You are in "linkedin.com", you cannot move forward any steps.
browserHistory.back(2);                   // You are in "linkedin.com", move back two steps to "facebook.com" then to "google.com". return "google.com"
browserHistory.back(7);                   // You are in "google.com", you can move back only one step to "leetcode.com". return "leetcode.com"
```

Constraints:

1 <= homepage.length <= 20  
1 <= url.length <= 20  
1 <= steps <= 100  
homepage and url consist of  '.' or lower case English letters.  
At most 5000 calls will be made to visit, back, and forward.  


## 方法


```go
type BrowserHistory struct {
  stack []string
  pos int
}


func Constructor(homepage string) BrowserHistory {
  return BrowserHistory{stack:[]string{homepage}}
}


func (this *BrowserHistory) Visit(url string)  {
  this.pos++
  this.stack = append(this.stack[:this.pos], url)
}


func (this *BrowserHistory) Back(steps int) string {
  newPos := this.pos - steps
  if newPos <= 0 {
    this.pos = 0
  } else {
    this.pos = newPos
  }
  return this.stack[this.pos]
}


func (this *BrowserHistory) Forward(steps int) string {
  newPos := this.pos + steps
  if newPos >= len(this.stack) {
    this.pos = len(this.stack) - 1
  } else {
    this.pos = newPos
  }
  return this.stack[this.pos]
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
 ```


 ```python
class BrowserHistory:

    def __init__(self, homepage: str):
        self.hist = [homepage]
        self.pos = 0

    def visit(self, url: str) -> None:
        self.pos += 1
        self.hist[self.pos:] = []
        self.hist.append(url)

    def back(self, steps: int) -> str:
        self.pos -= steps
        if self.pos < 0:
            self.pos = 0
        return self.hist[self.pos]

    def forward(self, steps: int) -> str:
        self.pos += steps
        if self.pos >= len(self.hist):
            self.pos = len(self.hist) - 1
        return self.hist[self.pos]



# Your BrowserHistory object will be instantiated and called as such:
# obj = BrowserHistory(homepage)
# obj.visit(url)
# param_2 = obj.back(steps)
# param_3 = obj.forward(steps)

 ```