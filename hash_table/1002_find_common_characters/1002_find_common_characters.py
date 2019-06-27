class Solution(object):
    def commonChars(self, A):
        """
        :type A: List[str]
        :rtype: List[str]
        """
        total_unique_chars = set(''.join(A[0]))
        result = []
        for char in total_unique_chars:
            char_count = []
            for i in A:
                char_count.append(i.count(char))
            result.extend([char]*min(char_count))  
        return result

A = ["bella","label","roller"]
print(Solution().commonChars(A))