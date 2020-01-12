
class Solution(object):
    def generateTrees(self, n):
        def generate_trees(start, end):
            if start > end:
                return [None,]
        
            all_trees = []
            for i in range(start, end +1):
                left_trees = generate_trees(start, i-1)
                right_trees = generate_trees(i+1, end)

                for l in left_trees:
                    for r in right_trees:
                        current_tree = TreeNode(i)
                        current_tree.left = l
                        current_tree.right = r
                        all_trees.append(current_tree)
            return all_trees
        return generate_trees(1, n) if n else []