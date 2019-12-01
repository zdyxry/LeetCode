
class Solution(object):
    def isValidSerialization(self, preorder):
        p = preorder.split(',')
        slot = 1
        for node in p:
            if slot == 0:
                return False
            if node == '#':
                slot -= 1
            else:
                slot += 1
        
        return slot == 0


preorder = "9,3,4,#,#,1,#,#,2,#,6,#,#"
res = Solution().isValidSerialization(preorder)
print(res)