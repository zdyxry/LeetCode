
# i < j < k 
# ai < ak < aj

# 下面这种方法利用来栈来做，既简洁又高效，思路是我们维护一个栈和一个变量third，
# 其中third就是第三个数字，也是pattern 132中的2，栈里面按顺序放所有大于third的数字，
# 也是pattern 132中的3，那么我们在遍历的时候，如果当前数字小于third，
# 即pattern 132中的1找到了，我们直接返回true即可，因为已经找到了，注意我们应该从后往前遍历数组。
# 如果当前数字大于栈顶元素，那么我们按顺序将栈顶数字取出，赋值给third，然后将该数字压入栈，
# 这样保证了栈里的元素仍然都是大于third的，我们想要的顺序依旧存在，进一步来说，
# 栈里存放的都是可以维持second > third的second值，其中的任何一个值都是大于当前的third值，
# 如果有更大的值进来，那就等于形成了一个更优的second > third的这样一个组合，
# 并且这时弹出的third值比以前的third值更大，为什么要保证third值更大，
# 因为这样才可以更容易的满足当前的值first比third值小这个条件。


class Solution(object):
    def find132pattern(self, nums):
        import sys
        stack = []
        s3 = -sys.maxint
        for n in nums[::-1]:
            if n < s3:
                return True
            while stack and stack[-1] < n:
                s3 = stack.pop()
            stack.append(n)
        return False


nums = [1,2,3,4]
res = Solution().find132pattern(nums)
print(res)
            