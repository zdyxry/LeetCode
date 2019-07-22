class MinStack(object):
    def __init__(self, s=[]):
        self.s = s
        self.t = []
        self.m = float('inf')

    def push(self, x):
        self.s.append(x)
        if x < self.m:
            self.m = x
        self.t.append(self.m)

    def pop(self):
        self.s.pop()
        self.t.pop()
        self.m = self.t[-1] if self.t else float('inf')

    def top(self):
        return self.s[-1] if self.s else None

    def getMin(self):
        return self.m if self.s else None


minStack = MinStack()
minStack.push(1)
minStack.push(2)
minStack.push(3)
print(minStack.top())
print(minStack.getMin())
minStack.pop()
print(minStack.top())
print(minStack.getMin())