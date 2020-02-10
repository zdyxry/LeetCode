
class Solution(object):
    def convert(self, s, numRows):
        value_list = [''] * numRows
        count = 0
        sign = 1
        res = ''

        if numRows == 1 or numRows >= len(s):
            res = s
            return res

        for char in s:
            print count, value_list
            value_list[count] += char
            if count == 0:
                sign = 1
            if count == numRows - 1:
                sign = -1
            
            count += sign
            
        
        res = ''.join(value_list)
        return res

s = "PAYPALISHIRING"
numRows = 3

res = Solution().convert(s, numRows)
print(res)