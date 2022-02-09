class Solution:
    def kthDistinct(self, arr: List[str], k: int) -> str:
        tmp = {}
        for i in arr:
            if i not in tmp:
                tmp[i] = 1
            else:
                tmp[i] += 1
        
        for i in arr:
            if tmp[i] == 1:
                k -= 1
            if k == 0:
                return i
        return ""
