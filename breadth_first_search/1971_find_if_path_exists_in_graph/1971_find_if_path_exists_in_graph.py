class Solution:
    def validPath(self, n: int, edges: List[List[int]], source: int, destination: int) -> bool:
        d = {}
        for i in edges:
            a, b  = i 
            d.setdefault(a, [])
            d.setdefault(b, [])
            d[a].append(b)
            d[b].append(a)
        
        queue = [source]
        vis = set()
        vis.add(source)
        while queue:
            h = queue.pop()
            if h == destination:
                return True
            for i in d[h]:
                if i not in vis:
                    queue.append(i)
                    vis.add(i)
        return False
