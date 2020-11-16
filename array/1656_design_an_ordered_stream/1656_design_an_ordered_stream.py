

class OrderedStream:

    def __init__(self, n: int):
        self.val = [0 for i in range(n+1)]
        self.ptr = 1


    def insert(self, id: int, value: str) -> List[str]:
        self.val[id] = value
        rtn_lst = []
        if id == self.ptr:
            while self.ptr < len(self.val) and self.val[self.ptr] != 0:
                rtn_lst.append(self.val[self.ptr])
                self.ptr += 1
        return rtn_lst



# Your OrderedStream object will be instantiated and called as such:
# obj = OrderedStream(n)
# param_1 = obj.insert(id,value)