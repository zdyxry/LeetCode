package main 

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