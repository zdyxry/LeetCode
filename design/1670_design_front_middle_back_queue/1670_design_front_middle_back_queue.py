

class FrontMiddleBackQueue:

    def __init__(self):
        self.q = []

    def pushFront(self, val: int) -> None:
        self.q.insert(0, val)

    def pushMiddle(self, val: int) -> None:
        self.q.insert(len(self.q) // 2, val)

    def pushBack(self, val: int) -> None:
        self.q.append(val)

    def popFront(self) -> int:
        return self.q.pop(0) if self.q else -1

    def popMiddle(self) -> int:
        n = len(self.q)
        if n % 2 == 0:
            n = n // 2 - 1
        else:
            n = n // 2
        return self.q.pop(n) if self.q else -1

    def popBack(self) -> int:
        return self.q.pop() if self.q else -1


# Your FrontMiddleBackQueue object will be instantiated and called as such:
# obj = FrontMiddleBackQueue()
# obj.pushFront(val)
# obj.pushMiddle(val)
# obj.pushBack(val)
# param_4 = obj.popFront()
# param_5 = obj.popMiddle()
# param_6 = obj.popBack()
