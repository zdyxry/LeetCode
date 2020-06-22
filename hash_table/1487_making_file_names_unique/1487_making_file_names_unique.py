from typing import List

class Solution:
    def getFolderNames(self, names: List[str]) -> List[str]:
        nameMap = {} # baseName : largest k suffix
        res = []
        
        for n in names:
            if n in nameMap:
                # find k
                k = nameMap[n] + 1
                while ( n + "(" + str(k) + ")" ) in nameMap:
                    k += 1
                nameMap[n] = k
                n = n + "(" + str(k) + ")" # with suffix is now considered a base name
                
            nameMap[n] = 0 # first time seeing this base name
            res.append(n)
        return res


names = ["pes","fifa","gta","pes(2019)"]
res = Solution().getFolderNames(names)
print(res)