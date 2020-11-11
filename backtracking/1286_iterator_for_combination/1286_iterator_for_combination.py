

class CombinationIterator:

    def __init__(self, characters: str, combinationLength: int):
        res=[]
        def helper(chars,length,path):
            if length==0:
                res.append(path[:])
            for index in range(len(chars)):
                path+=chars[index]
                helper(chars[index+1:],length-1,path)
                path=path[:-1]
        helper(characters,combinationLength,"")
        self.res=res
        self.res_len=len(res)
        
    def next(self) -> str:
        if self.hasNext():
            self.res_len-=1
            return self.res.pop(0)
        
    def hasNext(self) -> bool:
        if self.res_len==0:
            return False
        return True

# Your CombinationIterator object will be instantiated and called as such:
# obj = CombinationIterator(characters, combinationLength)
# param_1 = obj.next()
# param_2 = obj.hasNext()