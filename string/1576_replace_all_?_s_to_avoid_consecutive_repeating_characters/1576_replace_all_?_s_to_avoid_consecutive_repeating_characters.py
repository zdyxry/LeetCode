import string


class Solution:
    def modifyString(self, s: str) -> str:
        l = list(s)
        alphabet = list(string.ascii_lowercase)
        for i, _ in enumerate(l):
            if l[i] == "?":
                for j in alphabet:
                    if i == len(l) - 1:
                        if l[i-1] != j:
                            l[i] = j
                            break
                    elif l[i-1] != j and l[i+1] != j:
                        l[i] = j
                        break
        return ''.join(l)


s = "?zs"
res = Solution().modifyString(s)
print(res)