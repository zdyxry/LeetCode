from typing import List


class Solution:
    def decrypt(self, code: List[int], k: int) -> List[int]:
        result = [0] * len(code)
        if k > 0:
            for i in range(len(code)):
                tmp = 0
                for j in range(1, k+1):
                    tmp += code[(i+j)%len(code)]
                result[i] = tmp
            
        elif k < 0:
            for i in range(len(code)):
                tmp = 0
                for j in range(1, abs(k)+1):
                    tmp += code[(i-j)%len(code)]
                result[i] = tmp

        return result


code = [5,7,1,4]
k = 3
res = Solution().decrypt(code, k)
print(res)