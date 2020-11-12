from typing import List


class Solution:
    def findDuplicate(self, paths: List[str]) -> List[List[str]]:
        mp = {}
        for content in paths:
            path, *files = content.split(' ')
            for f in files:
                fn, content = f.split('(')
                content = content[:-1]
                if content not in mp:
                    mp[content] = []
                mp[content].append(f'{path}/{fn}')
        ret = []
        for k, v in mp.items():
            if len(v) > 1:
                ret.append(v)
        return ret



paths = ["root/a 1.txt(abcd) 2.txt(efgh)","root/c 3.txt(abcd)","root/c/d 4.txt(efgh)","root 4.txt(efgh)"]
res = Solution().findDuplicate(paths)
print(res)