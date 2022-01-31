class Solution:
    def findOriginalArray(self, changed: List[int]) -> List[int]:
        length = len(changed)
        changed.sort()
        if length % 2 != 0 :
            return []

        m = {}
        result = []
        for i in changed:
            if i %2 == 0  and i // 2 in m and m[i//2] > 0:
                m[i//2] -= 1
                result.append(i//2)
            else:
                if i not in m:
                    m[i] = 1
                else:
                    m[i] += 1
        for j in m.values():
            if j > 0 :
                return []
        return result