
class Solution:
    def interpret(self, command: str) -> str:
        return command.replace("()", "o").replace("(al)", "al")


command = "G()(al)"
res  = Solution().interpret(command)
print(res)