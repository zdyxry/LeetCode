class Solution:
    def finalValueAfterOperations(self, operations: List[str]) -> int:
        result = 0
        for o in operations:
            if '+' in o:
                result += 1
            if '-' in o:
                result -= 1

        return result