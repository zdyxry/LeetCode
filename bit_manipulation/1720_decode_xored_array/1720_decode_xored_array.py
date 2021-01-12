from typing import List

class Solution:
    def decode(self, encoded: List[int], first: int) -> List[int]:
        n = len(encoded)
            
        arr = [first] + [0] * (n)
        
        for i in range(n):
            arr[i + 1] = encoded[i] ^ arr[i]
        
        return arr


encode = [1,2,3]
first = 1
res = Solution().decode(encode, first)
print(res)