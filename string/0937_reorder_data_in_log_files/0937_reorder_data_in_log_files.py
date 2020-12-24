from typing import List

class Solution:
    def reorderLogFiles(self, logs: List[str]) -> List[str]:
        a=[]
        b=[]
        for i in logs:
            if i.split(" ")[1].isdigit():
                b.append(i)
            else:
                a.append(i)
        a=sorted(a,key=lambda a:(a.split(" ")[1:],a.split()[0]))
        return a+b


logs = ["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"]
res = Solution().reorderLogFiles(logs)
print(res)