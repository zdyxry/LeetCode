class Node(object):
    def __init__(self, val, children):
        self.val = val
        self.children = children

class Solution(object):
    def maxDepth(self, root):
        if not root:
            return 0
        queue = []
        queue.append(root)
        height = 0
        while len(queue) > 0:
            num_nodes = len(queue)
            for i in range(num_nodes):
                for child in queue[0].children:
                    queue.append(child)
                queue.remove(queue[0])
            height += 1
        return height