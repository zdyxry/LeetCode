class Solution:
    def minTime(self, n: int, edges: List[List[int]],
                hasApple: List[bool]) -> int:
        # 初始化路径
        maps = collections.defaultdict(list)
        for e in edges:
            maps[e[0]].append(e[1])

        def dfs(i):
            selfOrChildHasApple = hasApple[i]
            for nex in maps[i]:
                selfOrChildHasApple |= dfs(nex)
            if not selfOrChildHasApple:
                # 从字典中移除自身或子树都没有苹果的节点
                del maps[i]
            return selfOrChildHasApple

        dfs(0)
        # 字典个数即为最终有效节点的个数
        # 但是有可能有效节点为0, 所以需要max一下
        return max(0, 2 * (len(maps) - 1))



