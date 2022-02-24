class Solution:
    def capitalizeTitle(self, title: str) -> str:
        result = ""
        for i in title.split():
            if len(i) <= 2:
                result += i.lower()
            else:
                result += i.title()
            result += " "
        return result.strip()