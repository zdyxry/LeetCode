class Solution(object):
    def frequencySort(self, s):
        dic = {}
        result = ""

        for char in s:
            if char in dic:
                dic[char] += 1
            else:
                dic[char] =1
        
        sorted_dic = sorted(dic, key=dic.get, reverse=True)
        for count in sorted_dic:
            result += count * (dic[count])
        return result

s = "tree"
res = Solution().frequencySort(s)
print(res)