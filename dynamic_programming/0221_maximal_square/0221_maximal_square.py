class Solution(object):
    def maximalSquare(self, matrix):
        if not matrix or not matrix[0]:
            return 0
        rows, cols = len(matrix), len(matrix[0])
        max_side = 0
        square_sides = [0] * cols
        for r in range(rows):
            new_square_sides = [int(matrix[r][0])
                                ] + [0 for _ in range(cols - 1)]
            for c in range(1, cols):
                if matrix[r][c] == "1":
                    new_square_sides[c] = 1 + min(new_square_sides[c - 1],
                                                  square_sides[c],
                                                  square_sides[c - 1])
                    print(r, new_square_sides)
            max_side = max(max_side, max(new_square_sides))
            square_sides = new_square_sides
        return max_side**2


matrix = [["1", "0", "1", "0", "0"], ["1", "0", "1", "1", "1"],
          ["1", "1", "1", "1", "1"], ["1", "0", "0", "1", "0"]]
res = Solution().maximalSquare(matrix)
print(res)
