from typing import List


class Solution:
    def removeSubfolders(self, folder: List[str]) -> List[str]:
        res, t = [], ' '
        for f in sorted(folder):
            print(f)
            if not f.startswith(t):
                res.append(f)
                t = f + '/'
        return res


folder = ["/a","/a/b","/c/d","/c/d/e","/c/f"]
res = Solution().removeSubfolders(folder)
print(res)

